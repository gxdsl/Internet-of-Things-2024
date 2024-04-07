package main

import (
	"Server_Go/dataBase"
	"Server_Go/hardware"
	"Server_Go/routers"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net"
	"os"
)

func main() {
	//获取初始化的数据库
	dataBase.InitDB()
	//延迟关闭数据库
	defer func(DB *gorm.DB) {
		_ = DB.Close()
	}(dataBase.DB)

	// 创建第一个Gin路由,监听端口8888
	ginServer := gin.Default()
	//解决跨域问题
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	ginServer.Use(cors.New(config))
 
	// 运行路由组
	routers.GroupNestd(ginServer)

	// 设置服务器的 IP 地址和端口
	//serverIP := "0.0.0.0"
	serverIP := "127.0.0.1"
	serverPort := "8888"
	ShowIP() //打印端口
	// 监听端口，将其放在单独Go协程中运行
	go func() {
		ginErr := ginServer.Run(fmt.Sprintf("%s:%s", serverIP, serverPort))
		if ginErr != nil {
			fmt.Println("Gin Server启动失败:", ginErr)
			return
		}
	}()

	hardware.TCPServerInit()
}

func ShowIP() {
	serverPort := "8888"

	// 获取本机的主机名
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error getting hostname:", err)
		return
	}

	// 获取本机的 IP 地址
	ips, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Println("Error getting IP address:", err)
		return
	}

	// 打印 IP 地址
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			//goland:noinspection HttpUrlsUsage
			fmt.Printf("GinServer APP Runinng: http://%s:%s\n", ipv4, serverPort)
		}
	}
}
