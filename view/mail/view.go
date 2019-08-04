package mail

import (
	"fmt"
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
	from:=c.PostForm("from")
	content:=c.PostForm("content")
	eArr :=strings.Split(emails,",")
	fmt.Println(eArr,title,from,content)
	sql := `select status,info from hfish_setting where id = 1`
	isAlertStatus := dbUtil.Query(sql)
	info := isAlertStatus[0]["info"]
	config := strings.Split(info.(string), "&&")
	if from!=""{
		config[2]=from
	}
	send.SendMail(eArr,title,content,config)
	c.JSON(http.StatusOK,gin.H{
		"code":200,
		"msg":"success",
		"data":nil,
	})
}