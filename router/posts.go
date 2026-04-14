package router

import "github.com/gin-gonic/gin"

func regPostRoutes(r *gin.Engine) {
	posts := r.Group("/posts")

	posts.GET("/", nil)
	posts.GET("/:id", nil)
	posts.POST("/", nil)
	posts.PUT("/:id", nil)
	posts.DELETE("/:id", nil)
}
