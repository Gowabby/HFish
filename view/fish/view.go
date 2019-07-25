package fish

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Task List 页面
func Html(c *gin.Context) {
	c.HTML(http.StatusOK, "fish.html", gin.H{})
}