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
		fmt.Print(jsonData) //打印解析数据的切片

		// 根据Json数据类型执行不同操作
		switch {
		case jsonData["dispenser_id"] != nil && jsonData["tds"] != nil && jsonData["temperature"] != nil &&
			jsonData["flow"] != nil && jsonData["status"] != nil:
			WaterHandler(buffer.Bytes())
		case jsonData["dispenser_id"] != nil && jsonData["card"] != nil && jsonData["amount"] != nil:
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
	var jsonData struct {
		DispenserID uint `json:"dispenser_id"`
		Status      string
		TDS         float64
		Temperature float64
		Flow        bool
	}

	// 解析 JSON 数据
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Println("Error decoding water and temperature data:", err)
		return
	}

	// 创建 DispenserStatus 对象
	status := dataBase.DispenserStatus{
		DispenserID: jsonData.DispenserID,
		Status:      jsonData.Status,
		TDS:         jsonData.TDS,
		Temperature: jsonData.Temperature,
		Flow:        jsonData.Flow,
	}

	// 插入数据到数据库表中
	if err := dataBase.DB.Create(&status).Error; err != nil {
		fmt.Println("Error inserting data into database:", err)
		return
	}

	//打印插入的数据
	fmt.Println("Data inserted into database:", status.DispenserID, status.Status, status.Temperature, status.TDS, status.Flow)
}

// MoneyHandler 处理消费金额
func MoneyHandler(data []byte) {
	var jsonData map[string]interface{}

	// 解析 JSON 数据
	if err := json.Unmarshal(data, &jsonData); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// 获取饮水机ID和用水量
	dispenserID, ok1 := jsonData["dispenser_id"].(float64)
	amount, ok2 := jsonData["amount"].(float64)
	if !ok1 || !ok2 {
		fmt.Println("Invalid JSON format")
		return
	}

	// 根据饮水机ID查找饮水机的单价
	var waterDispenser dataBase.WaterDispenser
	if err := dataBase.DB.Where("id = ?", uint(dispenserID)).First(&waterDispenser).Error; err != nil {
		fmt.Println("Error querying water dispenser:", err)
		return
	}

	// 计算消费金额
	price := waterDispenser.Price
	totalAmount := amount * price

	// 获取用户卡号
	card, ok := jsonData["card"].(string)
	if !ok {
		fmt.Println("Invalid JSON format")
		return
	}

	// 根据卡号查找用户
	var user dataBase.User
	if err := dataBase.DB.Where("card = ?", card).First(&user).Error; err != nil {
		fmt.Println("Error querying user:", err)
		return
	}

	// 更新用户余额
	user.Balance -= totalAmount
	if err := dataBase.DB.Save(&user).Error; err != nil {
		fmt.Println("Error updating user balance:", err)
		return
	}

	// 插入消费记录
	transaction := dataBase.Transaction{
		User:        user.User,
		DispenserID: uint(dispenserID),
		Amount:      totalAmount,
	}
	if err := dataBase.DB.Create(&transaction).Error; err != nil {
		fmt.Println("Error inserting transaction:", err)
		return
	}

	fmt.Println("Transaction completed:", transaction)
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
