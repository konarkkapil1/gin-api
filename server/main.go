package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(gin *gin.Engine, address string) *http.Server {
	return &http.Server{
		Addr:    address,
		Handler: gin,
	}
}
