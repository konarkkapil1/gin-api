package router

import (
	"github.com/gin-gonic/gin"
	"github.com/konarkkapil1/gin-api/controllers"
	middleware "github.com/konarkkapil1/gin-api/middlewares"
)

func NewUserRouter(router gin.RouterGroup) {
	router.POST("/user", middleware.Authenticate, controllers.CreateUserController)
}
