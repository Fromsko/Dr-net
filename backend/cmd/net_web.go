package cmd

import (
	"guiforclash/backend/api"
	"guiforclash/backend/core"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func ScanApp(area string) *core.NetworkPair {
	// 启动
	result := core.StartApp(area)
	// 初始化对象
	NetPair := core.InitNetWorkPair()
	// 解析数据
	NetPair.ParserResult(result)
	// 存储
	NetPair.SaveJsonFile()

	if len(NetPair.Sect) != 0 {
		return NetPair
	}
	return nil
}

func GetRouter(result *core.NetworkPair) *gin.Engine {
	// 创建路由
	router = gin.Default()

	// 路由分组
	api.Router(router, result)

	return router
}
