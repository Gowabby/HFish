package mail

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"HFish/core/dbUtil"
	"HFish/utils/send"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "mail.html", gin.H{})
}

/*发送邮件*/
func SendEmailToUsers(c *gin.Context){
	emails:=c.PostForm("emails")
	title:=c.PostForm("title")
	content:=c.PostForm("content")
	eArr :=strings.Split(emails,",")
	sql := `select status,info from coot_setting where type=email;`
	isAlertStatus := dbUtil.Query(sql)
	send.SendMail(eArr,title,content,isAlertStatus)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":nil,
	})
}