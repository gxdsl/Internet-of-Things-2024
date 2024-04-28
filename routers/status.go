package routers

import (
	"Server_Go/dataBase"
	"Server_Go/hardware"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// OpenHandler 控制饮水机开水
func OpenHandler(ctx *gin.Context) {
	// 获取表单参数
	AppUser := ctx.PostForm("user")

	// 存储在全局变量中
	hardware.GlobalAppUser = AppUser

	// 数据验证
	if len(AppUser) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "用户名不能为空"})
		return
	}

	// 检查用户名是否存在
	var user dataBase.User
	if err := dataBase.DB.Where("user = ?", AppUser).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "用户名不存在",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "数据库查询失败",
			})
		}
		return
	}

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
