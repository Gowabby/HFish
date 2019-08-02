package view

import (
	"HFish/view/api"
	"HFish/view/dashboard"
	"HFish/view/fish"
	"HFish/view/mail"
	"HFish/view/setting"
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
	r.GET("/get/setting/info", setting.GetSettingInfo)
	r.POST("/post/setting/update",setting.UpdateEmailInfo)
	r.POST("/post/setting/login", setting.UpdateLoginInfo)
	r.POST("/post/setting/alertOver", setting.UpdateAlertOverInfo)
	r.POST("/post/setting/pushBullet", setting.UpdatePushBulletInfo)
	r.POST("/post/setting/pushFangTang", setting.UpdatePushFangTangInfo)
	r.POST("/post/setting/checkSetting",setting.UpdateStatusSetting)
	// API 接口
	r.GET("/api", api.Html)

	r.POST("/api/post/sendPageInfo",api.ReceivePageInfo)
	r.POST("/api/post/sendLoginInfo",api.ReceivePwdInfo)
}
