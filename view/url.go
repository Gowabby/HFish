package view

import (
	"HFish/view/dashboard"
	"HFish/view/mail"
	"HFish/view/setting"
	"HFish/view/fish"
	"HFish/view/api"
	"github.com/gin-gonic/gin"
	"HFish/view/login"
	"HFish/utils/conf"
)

func LoadUrl(r *gin.Engine) {
	// 登录
	r.GET("/login", login.Html)
	r.POST("/login", login.Login)
	r.GET("/logout", login.Logout)

	// 仪表盘
	r.GET("/", login.Jump, dashboard.Html)
	r.GET("/dashboard", login.Jump, dashboard.Html)
	r.GET("/get/dashboard/data", login.Jump, dashboard.GetFishData)

	// 钓鱼列表
	r.GET("/fish", login.Jump, fish.Html)
	r.GET("/get/fish/list", login.Jump, fish.GetFishList)
	r.GET("/get/fish/info", login.Jump, fish.GetFishInfo)
	r.POST("/post/fish/del", login.Jump, fish.PostFishDel)

	// 邮件群发
	r.GET("/mail", login.Jump, mail.Html)

	// 设置
	r.GET("/setting", login.Jump, setting.Html)

	// API 接口
	// WEB 上报钓鱼信息
	apiStatus := conf.Get("api", "status")

	// 判断 API 是否启用
	if apiStatus == "1" {
		apiUrl := conf.Get("api", "url")
		r.POST(apiUrl, api.ReportWeb)
	}
}
