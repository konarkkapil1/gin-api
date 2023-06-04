package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konarkkapil1/gin-api/services"
)

func CreateUserController(c *gin.Context) {
	user, err := services.CreateUserService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
