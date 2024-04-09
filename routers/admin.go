package routers

import (
	"Server_Go/dataBase"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// AdminLoginHandler 处理管理员登录请求
func AdminLoginHandler(ctx *gin.Context) {
	// 获取表单参数
	name := ctx.PostForm("user")
	password := ctx.PostForm("password")

	// 数据验证
	if len(name) == 0 || len(password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "用户名不能为空且密码长度不能少于6位"})
		return
	}

	// 检查用户名是否存在
	var admin dataBase.Admin
	if err := dataBase.DB.Where("username = ?", name).First(&admin).Error; err != nil {
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

	// 检查密码是否正确
	if admin.Password != password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "密码错误"})
		return
	}

	// 返回登录成功信息
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "登录成功",
	})
}

// AdminRegisterHandler 用户注册
func AdminRegisterHandler(ctx *gin.Context) {
	// 获取表单参数
	name := ctx.PostForm("user")
	password := ctx.PostForm("password")
	//card := ctx.PostForm("card") // 获取card参数

	// 数据验证
	if len(name) == 0 || len(password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户名不能为空，密码长度不能少于6位，卡号不能为空",
		})
		return
	}

	// 查询用户名和卡号是否已存在
	var existingAdmin dataBase.Admin
	if err := dataBase.DB.Where("username = ? ", name).First(&existingAdmin).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// 用户名和卡号都可用，创建新用户
			addAdmin := dataBase.Admin{
				Username: name,
				Password: password,
				Status:   true,
			}
			if err := dataBase.DB.Create(&addAdmin).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "新增管理员失败",
				})
				return
			}
			// 返回注册成功信息
			ctx.JSON(http.StatusCreated, gin.H{
				"code":    http.StatusCreated,
				"message": "注册成功",
			})
		} else {
			// 数据库查询错误
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "数据库查询失败",
			})
		}
		return
	}

	// 用户名或卡号已存在
	if existingAdmin.Username == name {
		ctx.JSON(http.StatusConflict, gin.H{
			"code":    409,
			"message": "用户名已存在",
		})
	}
}

// AdminListHandler 列出所有管理员
func AdminListHandler(ctx *gin.Context) {
	// 查询所有用户
	var users []dataBase.Admin
	result := dataBase.DB.Find(&users)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "查询用户失败",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "查询用户成功",
		"data":    users,
	})
}
