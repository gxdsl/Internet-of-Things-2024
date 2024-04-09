package routers

import (
	"github.com/gin-gonic/gin"
)

// GroupNestd 路由组
func GroupNestd(engine *gin.Engine) {
	Admin := engine.Group("/admin")
	{
		Admin.POST("/login", AdminLoginHandler)       //管理员登录
		Admin.POST("/register", AdminRegisterHandler) //新增管理员

		Admin.GET("/list", AdminListHandler) //列出所有管理员
	}
	User := engine.Group("/user")
	{
		User.POST("/login", LoginHandler)       //用户登录
		User.POST("/register", RegisterHandler) //用户注册
		User.GET("/list", ListHandler)          //列出所有用户
		User.POST("/delete", DeleteHandler)     //删除用户
		User.POST("/modify", ModifyHandler)     //修改密码
		User.POST("/recharge", RechargeHandler) //用户充值

		User.POST("/spendlatest", UserlatestHandler) //查询用户最新消费数据
		User.POST("/spendlist", UserlistHandler)     //查询用户最新20条消费数据
	}
	Data := engine.Group("/data")
	{
		Data.GET("/list", ChecklistHandler)       //查询最新20条状态数据
		Data.GET("/latest", ChecklatestHandler)   //查询最新一条状态数据
		Data.GET("/tdslatest", TDSlatestHandler)  //查询TDS水质最新数据
		Data.GET("/temlatest", TemplatestHandler) //查询水温最新数据

		Data.POST("/devicespend", DeviceslistHandler) //查询设备最新20条消费数据
	}
	Status := engine.Group("/status")
	{
		Status.POST("/open", OpenHandler)   //控制饮水机开水
		Status.POST("/close", CloseHandler) //控制饮水机关水
	}
}
