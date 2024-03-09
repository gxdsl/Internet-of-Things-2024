package routers

import (
	"Server_Go/hardware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// OpenHandler 控制饮水机开水
func OpenHandler(ctx *gin.Context) {
	// 构建 JSON 响应
	response := gin.H{
		"motor": 1,
	}

	// 返回 JSON 响应
	ctx.JSON(http.StatusOK, response)

	hardware.Send(response)
}

// CloseHandler 控制饮水机关水
func CloseHandler(ctx *gin.Context) {
	// 构建 JSON 响应
	response := gin.H{
		"motor": 0,
	}

	// 返回 JSON 响应
	ctx.JSON(http.StatusOK, response)

	hardware.Send(response)
}
