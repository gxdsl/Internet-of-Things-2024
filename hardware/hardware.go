package hardware

import (
	"encoding/json"
	"fmt"
	"net"
)

// Data 表示水质和温度数据
type Data struct {
	Tds         int `json:"tds"`
	Temperature int `json:"temperature"`
}

// User 表示用户数据
type User struct {
	UserID int `json:"userid"`
	Money  int `json:"money"`
}

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
	buffer := make([]byte, 2048)

	// 从连接中读取数据
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	// 解析 JSON 数据
	var jsonData map[string]interface{}
	if err := json.Unmarshal(buffer[:n], &jsonData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 根据数据类型执行不同操作
	switch {
	case jsonData["tds"] != nil && jsonData["temperature"] != nil:
		handleWaterAndTemperatureData(buffer[:n])
	case jsonData["userid"] != nil && jsonData["money"] != nil:
		handleUserData(buffer[:n])
	default:
		fmt.Println("Unknown data type")
	}
}

// 处理水质和温度数据
func handleWaterAndTemperatureData(data []byte) {
	var d Data
	if err := json.Unmarshal(data, &d); err != nil {
		fmt.Println("Error decoding water and temperature data:", err)
		return
	}

	// 在这里插入到数据库表data中
	fmt.Println("TDS:", d.Tds)
	fmt.Println("Temperature:", d.Temperature)
}

// 处理用户数据
func handleUserData(data []byte) {
	var u User
	if err := json.Unmarshal(data, &u); err != nil {
		fmt.Println("Error decoding user data:", err)
		return
	}

	// 在这里根据UserID找到对应的用户，并修改money列
	fmt.Println("UserID:", u.UserID)
	fmt.Println("Money:", u.Money)
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
