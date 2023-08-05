package api

import (
	"log"
	"net-tools/src/core"
	"path/filepath"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lifei6671/gocaptcha"
)

const (
	dx = 150
	dy = 50
)

func init() {
	fontDir := filepath.Join(core.ResDir, "fonts")
	if err := gocaptcha.ReadFonts(fontDir, ".ttf"); err != nil {
		panic(color.RedString("[DEBUG] 字体导入失败 "))
	}
}

// CaptchaHandler  验证码
func CaptchaHandler(c *gin.Context) {
	// 获取当前时间戳
	now := time.Now().Unix()

	// 获取时间戳参数
	timestamp := c.DefaultQuery("_", strconv.FormatInt(now, 10))

	// 获取上次请求的时间戳
	prev, exists := c.Get("timestamp")
	if exists {
		prevTimestamp, ok := prev.(int64)
		if ok {
			if now-prevTimestamp < 10 {
				// 时间间隔小于 10 秒，可能是请求过于频繁，进行处理
				c.JSON(400, gin.H{"error": "请求过于频繁，请稍后再试"})
				return
			}
		}
	}

	// 存储当前时间戳，供下次请求使用
	c.Set("timestamp", timestamp)

	captchaID, err := createCaptcha(c)
	if err != nil {
		log.Fatalln(err)
	}

	// 将验证码ID存储在Session中，用于校验
	session := sessions.Default(c)
	session.Set("captcha", captchaID)
	_ = session.Save()
}

// CreateCaptcha 创建一个验证码
func createCaptcha(c *gin.Context) (captchaID string, err error) {
	// 新的验证码
	captchaImage := gocaptcha.New(dx, dy, gocaptcha.RandLightColor())
	// 背景噪音
	captchaImage.DrawNoise(gocaptcha.CaptchaComplexLower)
	// 文字噪音
	captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexLower)
	// 验证码
	captchaID = gocaptcha.RandText(4)
	captchaImage.DrawText(captchaID)
	// captchaImage.Drawline(3);
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))
	captchaImage.DrawSineLine()
	//captchaImage.DrawHollowLine()

	// 返回验证码图片
	c.Header("Content-Type", "image/png")
	c.Header("Access-Control-Allow-Origin", "*")
	_ = captchaImage.SaveImage(c.Writer, gocaptcha.ImageFormatJpeg)

	return captchaID, nil
}
