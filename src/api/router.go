package api

import (
	"fmt"
	"net-tools/src/core"
	"net/http"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const PORT = 8080

var WebPair *core.NetworkPair

func init() {
	// 设置为生产环境
	gin.SetMode(gin.ReleaseMode)
}

func Router(router *gin.Engine, netPir *core.NetworkPair) {
	netPir.GetAreaList()
	fmt.Println(netPir.AreaList)
	fmt.Println(netPir.Sect)
	WebPair = netPir

	staticFiles(router)
	router.NoRoute(badRequestHandler)

	// 路由分组
	api := router.Group("/api/v1")
	{
		api.GET("/ip/search", randHandler)
		api.GET("/ip/check", checkHandler)
		api.GET("/ip/search/all_info", allHandler) // v1施工完成
	}

	// 基础路由
	router.GET("/", IndexHandler)
	router.POST("/login", LoginHandler)
	router.GET("/captcha", CaptchaHandler)

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

// 静态文件载入
func staticFiles(router *gin.Engine) {
	// 初始化 session 中间件
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	// 资源路径
	staticDir := filepath.Join(core.ResDir, "static")
	templatesDir := filepath.Join(core.ResDir, "templates/*")

	// 引入静态资源
	router.Static("/static", staticDir)
	router.LoadHTMLGlob(templatesDir)
}
