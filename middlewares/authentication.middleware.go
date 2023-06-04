package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	fmt.Println("Auth header: ", c.Request.Header["Authorization"])
	if c.Request.Header["Authorization"] == nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	}
	c.Next()
}