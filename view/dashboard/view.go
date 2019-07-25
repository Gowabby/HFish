package dashboard

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", gin.H{})
}
