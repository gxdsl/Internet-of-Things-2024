package routers

//func UpdatetimeHandler(ctx *gin.Context) {
//
//	// 从表单中获取要更新的数据字段
//	appTime := ctx.PostForm("time")
//
//	//// 将字段值转换为适当的数据类型
//	//temH, err := strconv.Atoi(temperatureH)
//	//if err != nil {
//	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "无法解析 TemH 字段"})
//	//	return
//	//}
//	//
//	//temL, err := strconv.Atoi(temperatureL)
//	//if err != nil {
//	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "无法解析 TemL 字段"})
//	//	return
//	//}
//	//
//	//TdsH, err := strconv.Atoi(sunTdsH)
//	//if err != nil {
//	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "无法解析 TdsH 字段"})
//	//	return
//	//}
//	//
//	ATime, err := strconv.Atoi(appTime)
//	if err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": "无法解析 appTime 字段"})
//		return
//	}
//
//	// 查询数据库中的特定记录（例如，id 为 1）
//	var status dataBase.Time
//	if err := dataBase.DB.First(&times, 1).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取指定记录"})
//		return
//	}
//
//	// 更新记录的字段值
//	Time.Time = ATime
//
//	// 更新数据库中的记录
//	if err := dataBase.DB.Save(&status).Error; err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新数据"})
//		return
//	}
//
//	// 返回成功响应
//	ctx.JSON(http.StatusOK, gin.H{"message": "数据已成功更新"})
//
//	jsondata := gin.H{
//		"Time": ATime,
//	}
//	hardware.Send(jsondata)
//}
