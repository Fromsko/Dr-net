package api

import (
	"fmt"
	"guiforclash/backend/core"
	"net/http"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

const PORT = 35203

var WebPair *core.NetworkPair

func init() {
	// 设置为生产环境
	gin.SetMode(gin.ReleaseMode)
}

func Router(router *gin.Engine, netPir *core.NetworkPair) {
	netPir.GetAreaList()
	WebPair = netPir
	router.NoRoute(badRequestHandler)

	// 路由分组
	api := router.Group("/api/v1")
	{
		api.GET("/ip/search", randHandler)
		api.GET("/ip/check", checkHandler)
		api.GET("/ip/search/all_info", allHandler) // v1施工完成
	}

	color.Green(fmt.Sprintf("[INFO] 服务器已启动\n访问地址 => http://localhost:%v", PORT))
	routerRegis(router)
}

// 获取注册路由信息
func routerRegis(router *gin.Engine) {
	// 获取注册的路径
	routes := router.Routes()
	for _, route := range routes {
		// 输出路径和HTTP方法
		color.HiBlue(fmt.Sprintf(
			"Router Path => http://localhost:%v%v  [ %v ]",
			PORT, route.Path, route.Method,
		))
	}
}

// 异常请求
func badRequestHandler(ctx *gin.Context) {
	statusCode := http.StatusBadRequest
	ctx.JSON(statusCode, gin.H{
		"status":  statusCode,
		"message": "Bad Request",
		"time":    core.NowTime(),
	})
}
