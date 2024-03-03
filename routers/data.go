package routers

import (
	"Server_Go/dataBase"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

// CheckAllHandler 查询所有数据
func CheckAllHandler(ctx *gin.Context) {
	var data []dataBase.Data

	// 使用 GORM 进行查询最新20条数据
	if err := dataBase.DB.Order("id desc").Limit(20).Find(&data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取数据",
		})
		return
	}

	//// 构建 JSON 响应
	//response := gin.H{
	//	"data": data,
	//}

	// 将 JSON 响应输出到客户端
	ctx.JSON(http.StatusOK, data)

}

// CheckLatestdataHandler 查询最新数据
func CheckLatestdataHandler(ctx *gin.Context) {
	var data dataBase.Data

	// 使用 GORM 进行查询最新数据
	if err := dataBase.DB.Order("id desc").Limit(1).Find(&data).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取数据",
		})
		return
	}

	// 将 JSON 响应输出到客户端
	ctx.JSON(http.StatusOK, data)
}

// SunCheckallHandler 查询所有水质TDS数据
func SunCheckallHandler(ctx *gin.Context) {

	// 查询单独一列 "tds" 数据
	var tdsValue []float64
	result := dataBase.DB.Table("data").Pluck("tds", &tdsValue)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tds": tdsValue})

}

func UploadAllHandler(ctx *gin.Context) {

	//获取JSON数据
	var uploaddata dataBase.Data
	if err := ctx.ShouldBindJSON(&uploaddata); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	////获取表单参数
	//tds := ctx.PostForm("tds")
	//temperature := ctx.PostForm("temperature")
	//personnel := ctx.PostForm("personnel")

	//数据验证
	//if uploaddata.tds == 0.0 {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code":    422,
	//		"message": "tds不为空",
	//	})
	//	return
	//}
	//if uploaddata.Temperature == 0.0 {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code":    422,
	//		"message": "Temperature不为空",
	//	})
	//	return
	//}
	//if len(uploaddata.Personnel) == 0 {
	//	ctx.JSON(http.StatusUnprocessableEntity, gin.H{
	//		"code":    422,
	//		"message": "Personnel不为空",
	//	})
	//	return
	//}

	//创建用户
	Upload := dataBase.Data{
		Tds:         uploaddata.Tds,
		Temperature: uploaddata.Temperature,
		//Personnel:   uploaddata.Personnel,
	}
	dataBase.DB.Create(&Upload)

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "上传成功",
	})
}

// SunChecklatestHandler 查询最新的水质TDS数据
func SunChecklatestHandler(ctx *gin.Context) {
	var tdsValue []float64

	// 查询最新10条 "tds" 数据
	result := dataBase.DB.Table("data").Select("tds").Order("id desc").Limit(1).Pluck("tds", &tdsValue)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tds": tdsValue[0]})
}

// TemChecklatestHandler 查询最新的温度数据
func TemChecklatestHandler(ctx *gin.Context) {
	var temperatureValue []float64

	// 查询最新10条 "tds" 数据
	result := dataBase.DB.Table("data").Select("temperature").Order("id desc").Limit(1).Pluck("temperature", &temperatureValue)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"temperature": int(math.Ceil(temperatureValue[0]))})
	//ctx.JSON(http.StatusOK, gin.H{"temperature": temperatureValue[0]})
}

//// UploadFileHandler 上传文件
//func UploadFileHandler(c *gin.Context) {
//	// 从请求中获取上传的文件
//	file, err := c.FormFile("file")
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	// 打开上传的文件
//	uploadedFile, err := file.Open()
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	defer func(uploadedFile multipart.File) {
//		err := uploadedFile.Close()
//		if err != nil {
//
//		}
//	}(uploadedFile)
//	// 读取文件内容
//	fileData, err := ioutil.ReadAll(uploadedFile)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//	// 将文件内容存入数据库（假设你有一个名为File的模型来表示文件）
//	newFile := dataBase.File{
//		Filename: file.Filename,
//		Content:  fileData,
//	}
//	if err := dataBase.DB.Create(&newFile).Error; err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{"message": "File uploaded and saved to  database successfully"})
//}

//func DownloadFileHandler(c *gin.Context) {
//
//	var file dataBase.File
//
//	// 按照时间戳字段降序排序，选择第一条记录（最新的数据）
//	if err := dataBase.DB.Order("created_at DESC").Last(&file).Error; err != nil {
//		// 处理数据库查询错误
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"error": "无法查询数据库",
//		})
//		return
//	}
//	c.JSON(http.StatusOK, file)
//	hardware.Send(file)
//
//}
