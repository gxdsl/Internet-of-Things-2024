package routers

import (
	"Server_Go/dataBase"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// LoginHandler 用户登录
func LoginHandler(ctx *gin.Context) {
	////获取JSON数据
	//var newuser dataBase.User
	//if err := ctx.ShouldBindJSON(&newuser); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//获取表单参数
	name := ctx.PostForm("user")
	password := ctx.PostForm("password")

	////数据验证
	//if len(name) == 0 {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code":    422,
	//		"message": "用户名不为空",
	//	})
	//	return
	//}
	//if len(password) < 6 {
	//	ctx.JSON(http.StatusNotAcceptable, gin.H{
	//		"code":    406,
	//		"message": "密码不能少于6位",
	//	})
	//	return
	//}

	// 数据验证
	if len(name) == 0 || len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "用户名不能为空且密码不能少于6位"})
		return
	}

	//判断用户名是否存在
	var user dataBase.User
	dataBase.DB.Where("user = ?", name).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "用户名不存在",
		})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "密码错误",
		})
		return
	}

	//返回结果，登陆成功
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
	})
}

// RegisterHandler 用户注册
func RegisterHandler(ctx *gin.Context) {

	////获取JSON数据
	//var newuser dataBase.User
	//if err := ctx.ShouldBindJSON(&newuser); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//获取表单参数
	name := ctx.PostForm("user")
	password := ctx.PostForm("password")

	////数据验证
	//if len(name) == 0 {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code":    422,
	//		"message": "用户名不能为空",
	//	})
	//	return
	//}
	//if len(password) < 6 {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code":    422,
	//		"message": "密码不能少于6位",
	//	})
	//	return
	//}

	// 数据验证
	if len(name) == 0 || len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "用户名不能为空且密码不能少于6位"})
		return
	}

	//判断用户名是否存在
	var user dataBase.User
	dataBase.DB.Where("user = ?", name).First(&user)
	if user.ID != 0 {
		ctx.JSON(http.StatusConflict, gin.H{
			"code":    409,
			"message": "用户名已存在",
		})
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}
	addUser := dataBase.User{
		User:     name,
		Password: string(hasedPassword),
	}
	dataBase.DB.Create(&addUser)

	//返回结果
	ctx.JSON(http.StatusCreated, gin.H{
		"code":    201,
		"message": "注册成功",
	})
}

// ListHandler 列出所有用户
func ListHandler(ctx *gin.Context) {
	// 查询所有用户
	var users []dataBase.User
	result := dataBase.DB.Find(&users)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// DeleteHandler 删除指定用户
func DeleteHandler(ctx *gin.Context) {

	////获取JSON数据
	//var newuser dataBase.User
	//if err := ctx.ShouldBindJSON(&newuser); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//获取删除用户ID
	idDelete := ctx.PostForm("id")

	// 查询用户是否存在
	var user dataBase.User
	result := dataBase.DB.First(&user, idDelete)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "未查询到用户",
		})
		return
	}

	// 删除用户
	deleteResult := dataBase.DB.Delete(&user)
	if deleteResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "删除用户错误",
		})
		return
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "删除成功",
	})
}

// ModifyHandler 修改用户密码
func ModifyHandler(ctx *gin.Context) {

	////获取JSON数据
	//var newuser dataBase.User
	//if err := ctx.ShouldBindJSON(&newuser); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}

	//获取参数
	name := ctx.PostForm("user")
	password := ctx.PostForm("password")

	////判断用户名是否存在
	//var user dataBase.User
	//dataBase.DB.Where("user = ?", name).First(&user)
	//if user.ID == 0 {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code":    422,
	//		"message": "用户名不存在",
	//	})
	//	fmt.Printf("用户名不存在")
	//	return
	//}

	//数据验证
	//if len(password) < 6 {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code":    422,
	//		"message": "密码不能少于6位",
	//	})
	//	return
	//}

	//数据验证
	if len(name) == 0 || len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "message": "用户名不能为空且密码不能少于6位"})
		return
	}

	//密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "密码加密错误",
		})
		return
	}

	// 保存更新后的用户信息
	updateResult := dataBase.DB.Model(&dataBase.User{}).Where("user = ?",
		name).Update("password", hasedPassword)
	if updateResult.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "密码更改错误",
		})
		return
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "密码更改成功",
	})
}
