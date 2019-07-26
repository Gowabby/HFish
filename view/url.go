package view

import (
	"HFish/view/dashboard"
	"HFish/view/mail"
	"HFish/view/setting"
	"HFish/view/fish"
	"HFish/view/api"
	"github.com/gin-gonic/gin"
)

func LoadUrl(r *gin.Engine) {
	r.GET("/", dashboard.Html)
	// 仪表盘
	r.GET("/dashboard", dashboard.Html)
	// 钓鱼列表
	r.GET("/fish", fish.Html)
	// 邮件群发
	r.GET("/mail", mail.Html)
	// 设置
	r.GET("/setting", setting.Html)
	// API 接口
	r.GET("/api", api.Html)

	r.POST("/api/post/sendPageInfo",api.ReceivePageInfo)
	r.POST("/api/post/sendLoginInfo",api.ReceivePwdInfo)
}
