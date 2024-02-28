package hardware

import (
	"Server_Go/dataBase"
	"encoding/json"
	"fmt"
	"net"
)

var globalConn net.Conn

func Init() {
	// 创建第二个Gin引擎实例，监听端口9999
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			fmt.Println("Error Close:", err)
		}
	}(listener)
	fmt.Println("Server is listening on :9999")
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

// HandleClient 接受硬件
func HandleClient(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// 创建一个缓冲区来存储接收到的数据
	buffer := make([]byte, 2048)
	for {
		// 从连接中读取数据
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading data:", err)
			return
		}

		// 处理接收到的 JSON 数据
		data := buffer[:n]

		// 解析 JSON 数据
		var jsonData map[string]interface{}
		if err := json.Unmarshal(data, &jsonData); err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}

		//{
		//	"tds": 25,
		//	"temperature": 14,
		//	"personnel": "1",
		//	"created_time": "2023-10-1 1:2:11",
		//	"people": 12
		//}

		// 创建一个新的 Data 结构并填充字段
		newData := dataBase.Data{
			Tds:         jsonData["tds"].(float64),
			Temperature: jsonData["temperature"].(float64),
			//Personnel:   jsonData["personnel"].(string),
			CreatedTime: jsonData["created_time"].(string),
			//People:      jsonData["people"].(uint8),
			//CreatedTime: jsonData["created_time"].(datetime),
		}

		// 将字段转换为 uint16
		//peopleFloat := jsonData["people"].(float64)
		//dayFloat := jsonData["day"].(float64)
		//newData.People = uint16(peopleFloat)
		//newData.Day = uint16(dayFloat)
		//aFloat := jsonData["a"].(float64)
		//newData.A = uint16(aFloat)
		//
		//bFloat := jsonData["b"].(float64)
		//newData.B = uint16(bFloat)

		//cFloat := jsonData["c"].(float64)
		//newData.C = uint16(cFloat)
		//
		//dFloat := jsonData["d"].(float64)
		//newData.D = uint16(dFloat)

		//cFloat := jsonData["Time"].(float64)
		//newData.C = uint16(cFloat)
		//
		//dFloat := jsonData["d"].(float64)
		//newData.D = uint16(dFloat)

		// 插入数据到数据库表中
		if err := dataBase.DB.Create(&newData).Error; err != nil {
			fmt.Println("Error inserting data into database:", err)
			return
		}

		fmt.Println("Data inserted into database:", newData.ID, newData.Temperature, newData.Tds,
			newData.CreatedTime)

		//personstatus := jsonData["personnel"].(string)
		//ReactionHandle(personstatus)
	}
}

//func ReactionHandle(personnel string) {
//	db := dataBase.DB
//	var status dataBase.Status
//	if personnel == "1" {
//		db.Model(&status).Update("lamp", 1)
//		jsondata := gin.H{
//			"loud": 1,
//			"lamp": 1,
//		}
//		Send(jsondata)
//	}
//}

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

func Sendaudio(File interface{}) {
	jsonaudio, err := json.Marshal(File)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	type audio struct {
	}
	// 发送 JSON 字符串到客户端连接
	_, err = globalConn.Write(jsonaudio)
	if err != nil {
		fmt.Println("Error sending JSON data:", err)
		return
	}
}
