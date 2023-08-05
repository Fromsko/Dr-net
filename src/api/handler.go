package api

import (
	"math/rand"
	"net-tools/src/core"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// randHandler 获取随机数据
func randHandler(ctx *gin.Context) {
	// 获取随机列表
	rand.New(rand.NewSource(time.Now().Unix()))
	areaList := WebPair.AreaList

	randIndex := rand.Intn(len(areaList))
	randArea := WebPair.Sect[areaList[randIndex]]

	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "1"))

	if limit == 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"content": core.GetRandData(randArea, 1)[0],
			"limit":   limit,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"content": core.GetRandData(randArea, limit),
			"limit":   limit,
		})
	}
}

// allHandler 获取全部数据
func allHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, WebPair)
}

// checkHandler 检测网络连接状况
func checkHandler(ctx *gin.Context) {
	jsonResponse := gin.H{
		"status": http.StatusOK,
		"error":  "",
	}
	var parserContent, err = LoginCheck()
	if err != nil {
		jsonResponse["error"] = "获取登陆信息失败"
		ctx.JSON(http.StatusOK, jsonResponse)
		return
	}
	jsonResponse["data"] = parserContent
	ctx.JSON(http.StatusOK, jsonResponse)
}
