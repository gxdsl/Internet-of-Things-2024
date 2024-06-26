package routers

import (
	"Server_Go/dataBase"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"math"
	"net/http"
	"strconv"
)

// ChecklistHandler 查询数据
func ChecklistHandler(ctx *gin.Context) {
	var allData []dataBase.DispenserStatus

	// 从最新的数据开始，每隔300条获取一次，获取十次数据
	for i := 0; i < 10; i++ {
		offset := i * 300 // 计算偏移量
		var data dataBase.DispenserStatus

		if err := dataBase.DB.Order("id desc").Offset(offset).Limit(1).Find(&data).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "无法获取数据",
			})
			return
		}

		// 将获取的数据添加到所有数据中
		allData = append(allData, data)
	}

	// 将所有数据倒序发送到客户端
	reversedData := make([]dataBase.DispenserStatus, len(allData))
	for i, j := 0, len(allData)-1; i < len(allData); i, j = i+1, j-1 {
		reversedData[i] = allData[j]
	}

	ctx.JSON(http.StatusOK, reversedData)
}

// ChecklatestHandler 查询最新数据
func ChecklatestHandler(ctx *gin.Context) {
	var data dataBase.DispenserStatus

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

// TDSlatestHandler 查询最新的水质TDS数据
func TDSlatestHandler(ctx *gin.Context) {
	var tdsValue []float64

	// 查询最新10条 "tds" 数据
	result := dataBase.DB.Table("dispenser_statuses").Select("tds").Order("id desc").
		Limit(1).Pluck("tds", &tdsValue)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tds": tdsValue[0]})
}

// TemplatestHandler 查询最新的温度数据
func TemplatestHandler(ctx *gin.Context) {
	var temperatureValue []float64

	// 查询最新10条 "tds" 数据
	result := dataBase.DB.Table("dispenser_statuses").Select("temperature").Order("id desc").
		Limit(1).Pluck("temperature", &temperatureValue)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"temperature": int(math.Ceil(temperatureValue[0]))})
	//ctx.JSON(http.StatusOK, gin.H{"temperature": temperatureValue[0]})
}

// DeviceslistHandler 查询设备最新20条消费记录
func DeviceslistHandler(ctx *gin.Context) {
	// 从表单数据中获取设备ID
	deviceID := ctx.PostForm("device")

	// 数据验证
	if len(deviceID) == 0 {
		// 如果设备ID为空，返回400 Bad Request状态码和错误消息
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "设备ID不能为空",
		})
		return
	}

	// 查询设备是否存在
	if err := dataBase.DB.Where("id = ?", deviceID).First(&dataBase.WaterDispenser{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 如果设备不存在，返回404 Not Found状态码和错误消息
			ctx.JSON(http.StatusNotFound, gin.H{
				"code":    http.StatusNotFound,
				"message": fmt.Sprintf("未找到设备 '%s'", deviceID),
			})
			return
		}
		// 如果查询设备存在时出错，返回500 Internal Server Error状态码和错误消息
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "查询设备时出错",
		})
		return
	}

	// 查询设备最新的20条消费记录，按照ID字段的倒序排列
	var latestTransactions []dataBase.Transaction
	if err := dataBase.DB.Where("dispenser_id = ?", deviceID).Order("id desc").
		Limit(20).Find(&latestTransactions).Error; err != nil {
		// 如果查询设备消费记录时出错，返回500 Internal Server Error状态码和错误消息
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "查询设备最新20条消费记录时出错",
		})
		return
	}

	// 返回查询结果
	ctx.JSON(http.StatusOK, latestTransactions)
}

// SpendtotalHandler 查询设备消费记录总数
func SpendtotalHandler(ctx *gin.Context) {
	var total int64
	if err := dataBase.DB.Model(&dataBase.Transaction{}).Count(&total).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Total consumption acquisition failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"total": total,
	})
}

// SpenddataHandler 返回当前页的消费数据，按倒序排列
func SpenddataHandler(ctx *gin.Context) {
	// 从表单中获取页码和每页数量，默认为第一页，每页显示 10 条数据
	page, _ := strconv.Atoi(ctx.PostForm("page"))
	pageSize, _ := strconv.Atoi(ctx.PostForm("pageSize"))

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 查询当前页的消费数据，按ID倒序排列
	var transactions []dataBase.Transaction
	if err := dataBase.DB.Order("id desc").Offset(offset).Limit(pageSize).Find(&transactions).Error; err != nil {
		log.Println("Error getting spend data:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get spend data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"pagedata": transactions,
	})
}
