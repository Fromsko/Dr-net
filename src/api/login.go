package api

import (
	"time"

	"github.com/gin-gonic/gin"
)

// IndexHandler 主页
func IndexHandler(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

// LoginHandler 登陆函数
func LoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	captcha := c.PostForm("captcha")
	AuthLogin(c, username, password, captcha)
}

// AuthLogin 登陆认证
func AuthLogin(engine *gin.Context, user, pswd, captch string) {
	// 假设用户名密码错误
	if user != "admin" || pswd != "admin" {
		engine.JSON(400, gin.H{
			"error": map[string]any{
				"type":    "username_password_error",
				"message": "用户名或密码错误",
			},
		})
		return
	}

	engine.JSON(200, map[string]any{
		"code":    200,
		"message": "login success",
		"time":    time.Now(),
	})
}
