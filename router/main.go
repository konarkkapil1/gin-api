package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	routes := router.Group("api/v1/")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	NewUserRouter(*routes)
}
