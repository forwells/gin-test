package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func regTestRoutes(r *gin.Engine) {
	test := r.Group("/test")

	test.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "你好 gin 1112!"})
	})
}
