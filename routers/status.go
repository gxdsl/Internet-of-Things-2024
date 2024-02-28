package routers

import (
	"Server_Go/dataBase"
	"Server_Go/hardware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetLampHandler 读取小灯状态
func GetLampHandler(ctx *gin.Context) {
	var status dataBase.Status

	// 使用 GORM 查询第一行数据
	if err := dataBase.DB.First(&status).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取 lamp 值",
		})
		return
	}

	// 提取 lamp 字段的值
	lampValue := status.Lamp

	// 构建 JSON 响应
	response := gin.H{
		"lamp": lampValue,
	}

	// 返回 JSON 响应
	ctx.JSON(http.StatusOK, response)
}

// ToggleLampHandler 切换小灯状态
func ToggleLampHandler(ctx *gin.Context) {
	var status dataBase.Status

	// 查询当前 lamp 值
	if err := dataBase.DB.First(&status).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取 lamp 值",
		})
		return
	}

	// 切换 lamp 值 (0 到 1 或 1 到 0)
	newLampValue := 1 - status.Lamp

	// 更新数据库中的 lamp 值
	if err := dataBase.DB.Model(&status).Update("Lamp", newLampValue).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法切换 lamp 值",
		})
		return
	}

	// 构建 JSON 响应
	response := gin.H{
		"lamp": newLampValue,
	}

	// 返回 JSON 响应
	ctx.JSON(http.StatusOK, response)

	hardware.Send(response)
}

func UpdatedateHandler(ctx *gin.Context) {

	// 从表单中获取要更新的数据字段
	temperatureH := ctx.PostForm("temh")
	temperatureL := ctx.PostForm("teml")
	sunlightH := ctx.PostForm("lighth")
	sunlightL := ctx.PostForm("lightl")
	time := ctx.PostForm("time")
	st := ctx.PostForm("st")

	//// 将字段值转换为适当的数据类型
	//temH, err := strconv.Atoi(temperatureH)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "无法解析 TemH 字段"})
	//	return
	//}
	//
	//temL, err := strconv.Atoi(temperatureL)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "无法解析 TemL 字段"})
	//	return
	//}
	//
	//lightH, err := strconv.Atoi(sunlightH)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "无法解析 LightH 字段"})
	//	return
	//}
	//
	//lightL, err := strconv.Atoi(sunlightL)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "无法解析 LightL 字段"})
	//	return
	//}
	//time, err := strconv.Atoi(Time)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "无法解析 Time 字段"})
	//	return
	//}
	//st, err := strconv.Atoi(ST)
	//if err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "无法解析 ST 字段"})
	//	return
	//}

	// 查询数据库中的特定记录（例如，id 为 1）
	var status dataBase.Status
	if err := dataBase.DB.First(&status, 1).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取指定记录"})
		return
	}

	// 更新记录的字段值
	status.TemH = temperatureH
	status.TemL = temperatureL
	status.LightH = sunlightH
	status.LightL = sunlightL
	status.Time = time
	status.ST = st

	// 更新数据库中的记录
	if err := dataBase.DB.Save(&status).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新数据"})
		return
	}

	// 返回成功响应
	ctx.JSON(http.StatusOK, gin.H{"message": "数据已成功更新"})

	jsondata := gin.H{
		"TemH":   temperatureH,
		"TemL":   temperatureL,
		"LightH": sunlightH,
		"LightL": sunlightL,
		"Tiem":   time,
		"ST":     st,
	}

	hardware.Send(jsondata)
}

func GetdataHandler(ctx *gin.Context) {
	var status dataBase.Status

	// 使用 GORM 查询第一行数据
	if err := dataBase.DB.First(&status).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "无法获取值",
		})
		return
	}

	// 提取字段的值
	LightLValue := status.LightL
	LightHValue := status.LightH
	TemLValue := status.TemL
	TemHValue := status.TemH
	TimeValue := status.Time
	STValue := status.ST

	// 构建 JSON 响应
	response := gin.H{
		"lightl": LightLValue,
		"lighth": LightHValue,
		"temh":   TemHValue,
		"teml":   TemLValue,
		"time":   TimeValue,
		"st":     STValue,
	}

	// 返回 JSON 响应
	ctx.JSON(http.StatusOK, response)
}

//// UpdatepeopleHandler 更新人员
//func UpdatepeopleHandler(ctx *gin.Context) {
//
//	var status dataBase.Status
//
//	// 从表单中获取要更新的数据字段
//	people := ctx.PostForm("people")
//
//	// 查询要更新的记录（假设要更新的记录的ID为1）
//	if err := dataBase.DB.First(&status, 1).Error; err != nil {
//		if errors.Is(err, gorm.ErrRecordNotFound) {
//			ctx.JSON(http.StatusNotFound, gin.H{
//				"error": "记录未找到",
//			})
//		} else {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": "无法获取记录",
//			})
//		}
//		return
//	}
//
//	if people == "1" {
//
//		// 更新 People 字段的值加一
//		status.People++
//
//		// 执行更新操作
//		if err := dataBase.DB.Save(&status).Error; err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": "无法更新记录",
//			})
//			return
//		}
//
//		// 成功更新，返回成功响应
//		ctx.JSON(http.StatusOK, gin.H{
//			"message": "人数已增加",
//		})
//	}
//	if people == "2" {
//
//		// 更新 People 字段的值减一
//		status.People--
//
//		// 执行更新操作
//		if err := dataBase.DB.Save(&status).Error; err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": "无法更新记录",
//			})
//			return
//		}
//
//		// 成功更新，返回成功响应
//		ctx.JSON(http.StatusOK, gin.H{
//			"message": "人数已减少",
//		})
//	}
//	if people == "0" {
//		// 更新 People 字段的值清零
//		status.People = 0
//
//		// 执行更新操作
//		if err := dataBase.DB.Save(&status).Error; err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{
//				"error": "无法更新记录",
//			})
//			return
//		}
//
//		// 成功更新，返回成功响应
//		ctx.JSON(http.StatusOK, gin.H{
//			"message": "人数已清零",
//		})
//	}
//}
