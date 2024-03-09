package routers

import (
	"Server_Go/dataBase"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

// LoginHandler 处理用户登录请求
func LoginHandler(ctx *gin.Context) {
	// 获取表单参数
	name := ctx.PostForm("user")
	password := ctx.PostForm("password")

	// 数据验证
	if len(name) == 0 || len(password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "用户名不能为空且密码长度不能少于6位"})
		return
	}

	// 检查用户名是否存在
	var user dataBase.User
	if err := dataBase.DB.Where("user = ?", name).First(&user).Error; err != nil {
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
	if user.Password != password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "密码错误"})
		return
	}

	// 返回登录成功信息
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "登录成功",
	})
}

// RegisterHandler 用户注册
func RegisterHandler(ctx *gin.Context) {
	// 获取表单参数
	name := ctx.PostForm("user")
	password := ctx.PostForm("password")
	card := ctx.PostForm("card") // 获取card参数

	// 数据验证
	if len(name) == 0 || len(password) < 6 || len(card) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户名不能为空，密码长度不能少于6位，卡号不能为空",
		})
		return
	}

	// 查询用户名和卡号是否已存在
	var existingUser dataBase.User
	if err := dataBase.DB.Where("user = ? OR card = ?", name, card).First(&existingUser).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			// 用户名和卡号都可用，创建新用户
			addUser := dataBase.User{
				User:     name,
				Password: password,
				Card:     card, // 保存card参数
			}
			if err := dataBase.DB.Create(&addUser).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "用户注册失败",
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
	if existingUser.User == name {
		ctx.JSON(http.StatusConflict, gin.H{
			"code":    409,
			"message": "用户名已存在",
		})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{
			"code":    409,
			"message": "卡号已存在",
		})
	}
}

//// ListHandler 列出所有用户
//func ListHandler(ctx *gin.Context) {
//	// 查询所有用户
//	var users []dataBase.User
//	result := dataBase.DB.Find(&users)
//	if result.Error != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
//		return
//	}
//
//	ctx.JSON(http.StatusOK, users)
//}

// ListHandler 列出所有用户
func ListHandler(ctx *gin.Context) {
	// 查询所有用户
	var users []dataBase.User
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

// DeleteHandler 通过用户名删除用户
func DeleteHandler(ctx *gin.Context) {
	// 绑定表单数据到结构体
	var form struct {
		Username string `form:"user"`
	}
	if err := ctx.ShouldBind(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "无法解析表单数据",
		})
		return
	}

	// 判断用户名是否为空
	if form.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户名不能为空",
		})
		return
	}

	// 在数据库中查找要删除的用户
	var user dataBase.User
	result := dataBase.DB.Where("user = ?", form.Username).First(&user)
	if result.Error != nil {
		if gorm.IsRecordNotFoundError(result.Error) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "用户不存在",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "无法删除用户",
			})
		}
		return
	}

	// 删除用户
	result = dataBase.DB.Delete(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "无法删除用户",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "用户删除成功",
	})
}

// / ModifyHandler 更新用户密码
func ModifyHandler(ctx *gin.Context) {
	// 获取表单参数
	username := ctx.PostForm("user")
	newPassword := ctx.PostForm("password")

	// 数据验证
	if len(username) == 0 || len(newPassword) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "用户名不能为空，新密码长度不能少于6位",
		})
		return
	}

	// 查询用户
	var user dataBase.User
	if err := dataBase.DB.Where("user = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": "用户不存在",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code":    http.StatusInternalServerError,
				"message": "数据库查询失败",
			})
		}
		return
	}

	// 更新用户密码
	user.Password = newPassword
	if err := dataBase.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "密码更新失败",
		})
		return
	}

	// 返回密码更新成功信息
	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "密码更新成功",
	})
}
