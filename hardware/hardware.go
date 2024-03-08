package hardware

import (
	"Server_Go/dataBase"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

var globalConn net.Conn

// TCPServerInit 初始化TCP服务器
func TCPServerInit() {
	// 创建TCP服务器，监听端口9999
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()
	fmt.Println("TCPServer is listening on:9999")

	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		globalConn = conn
		go HandleClient(globalConn)
	}
}

// HandleClient 处理客户端连接
func HandleClient(conn net.Conn) {
	defer conn.Close()

	// 创建一个缓冲区来存储接收到的数据
	var buffer bytes.Buffer

	for {
		// 从连接中读取数据
		tmpBuffer := make([]byte, 2048)
		n, err := conn.Read(tmpBuffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading data:", err)
			}
			break // 退出循环
		}

		// 将读取到的数据追加到缓冲区中
		buffer.Write(tmpBuffer[:n])

		// 尝试解析 JSON 数据
		var jsonData map[string]interface{}
		if err := json.Unmarshal(buffer.Bytes(), &jsonData); err != nil {
			fmt.Println("Error decoding JSON:", err)
			buffer.Reset() // 清空缓冲区
			continue       // 继续等待数据
		}

		// 根据Json数据类型执行不同操作
		switch {
		case jsonData["tds"] != nil && jsonData["temperature"] != nil:
			WaterHandler(buffer.Bytes())
		case jsonData["cardid"] != nil && jsonData["money"] != nil:
			MoneyHandler(buffer.Bytes())
		default:
			fmt.Println("Unknown data type")
		}

		// 从缓冲区中移除已处理的数据
		buffer.Next(len(buffer.Bytes()))
	}
}

// WaterHandler 处理水质和温度数据
func WaterHandler(data []byte) {
	var jsonData dataBase.Data

	// 解析 JSON 数据到 Data 结构
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Println("Error decoding water and temperature data:", err)
		return
	}

	// 插入数据到数据库表中
	if err := dataBase.DB.Create(&jsonData).Error; err != nil {
		fmt.Println("Error inserting data into database:", err)
		return
	}

	fmt.Println("Data inserted into database:", jsonData.Temperature, jsonData.Tds)
}

// MoneyHandler 处理用户卡余额消费
func MoneyHandler(data []byte) {
	var jsonData dataBase.User

	// 解析 JSON 数据到 User 结构
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Println("解码用户数据出错:", err)
		return
	}

	// 在这里根据 CardID 找到对应的用户，并修改 Money 列
	var user dataBase.User
	if err := dataBase.DB.Where("card_id = ?", jsonData.CardID).First(&user).Error; err != nil {
		fmt.Println("查找用户出错:", err)
		return
	}

	// 更新用户的 Money 列
	user.Money = jsonData.Money
	if err := dataBase.DB.Save(&user).Error; err != nil {
		fmt.Println("更新用户余额出错:", err)
		return
	}

	// 打印更改完成的信息
	fmt.Printf("%s，卡号为 %s 的用户 %s 的余额为 %f\n", user.UpdatedAt, user.CardID, user.User, user.Money)
}

func Send(Data interface{}) {
	// 序列化 JSON 对象为 JSON 字符串
	jsonData, err := json.Marshal(Data)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// 发送 JSON 字符串到客户端连接
	_, err = globalConn.Write(jsonData)
	if err != nil {
		fmt.Println("Error sending JSON data:", err)
		return
	}
}
