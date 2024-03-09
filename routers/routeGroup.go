package routers

import (
	"github.com/gin-gonic/gin"
)

// GroupNestd 路由组
func GroupNestd(engine *gin.Engine) {
	User := engine.Group("/user")
	{
		User.POST("/login", LoginHandler)
		User.POST("/register", RegisterHandler)
		User.POST("/list", ListHandler)
		User.POST("/delete", DeleteHandler)
		User.POST("/modify", ModifyHandler)
	}
	//Data := engine.Group("/data")
	//{
	//	Upload := Data.Group("/upload")
	//	{
	//		Upload.POST("/all", UploadAllHandler)
	//		//Upload.POST("/file", UploadFileHandler)
	//	}
	//	Check := Data.Group("/check")
	//	{
	//		Check.GET("/alldata", CheckAllHandler)
	//		//Check.POST("/alldata", CheckAllHandler)
	//		Check.GET("/latestdata", CheckLatestdataHandler)
	//		Check.POST("/latestdata", CheckLatestdataHandler)
	//		Check.POST("/tdsall", SunCheckallHandler)
	//		Check.POST("/tdslatest", SunChecklatestHandler)
	//		Check.POST("/temlitlatest", TemChecklatestHandler)
	//	}
	//	//Download := Data.Group("/download")
	//	//{
	//	//	Download.POST("/file", DownloadFileHandler)		//上传音频文件
	//	//}
	//}
	//Status := engine.Group("/status")
	//{
	//	Status.GET("/getlamp", GetLampHandler)
	//	Status.POST("/togglelamp", ToggleLampHandler)
	//	Status.GET("/getdata", GetdataHandler)
	//	Status.POST("/updatedata", UpdatedateHandler)
	//	//Status.POST("/updatetime", UpdatetimeHandler)
	//
	//	//Status.POST("/updatepeople", UpdatepeopleHandler)
	//
	//}
}
