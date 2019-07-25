package mail

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "mail.html", gin.H{})
}
