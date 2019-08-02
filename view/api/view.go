package api

import (
	"github.com/gin-gonic/gin"
	"HFish/core/report"
	"net/http"
	"HFish/error"
)

func ReportWeb(c *gin.Context) {
	name := c.PostForm("name")
	info := c.PostForm("info")
	ip := c.ClientIP()

	report.ReportWeb(name, ip, info)

	c.JSON(http.StatusOK, error.ErrSuccessNull())
}
