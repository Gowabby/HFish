package api

import (
	"HFish/core/dbUtil"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.web", gin.H{})
}


/*记录访问者基本信息*/
func ReceivePageInfo(c *gin.Context){
	ip:=c.PostForm("ip")
	city:=c.PostForm("city")
	browser:=c.PostForm("browser")
	version:=c.PostForm("version")
	os:=c.PostForm("os")
	pageType:=c.PostForm("pageType")
	sql := `
		INSERT INTO user_info (
			ip,
			city,
			browser,
			version,
			os,
			pageType,
			create_time
		)
		VALUES
			(?,?,?,?,?,?,?);
		`
	dbUtil.Insert(sql, ip, city,browser,version, os,pageType,time.Now().Format("2006-01-02 15:04:05"))
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":nil,
	})
}
/*接收账号密码信息*/
func ReceivePwdInfo(c *gin.Context){
	loginName:=c.PostForm("loginName")
	loginPwd:=c.PostForm("loginPwd")
	pageType:=c.PostForm("pageType")
	sql := `
		INSERT INTO user_info (
			loginName,
			loginPwd,
			pageType,
			create_time
		)
		VALUES
			(?,?,?,?);
		`
	dbUtil.Insert(sql, loginName,loginPwd,pageType,time.Now().Format("2006-01-02 15:04:05"))
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":nil,
	})
}