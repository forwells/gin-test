package router

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	r := gin.Default()

	regTestRoutes(r)

	return r
}
