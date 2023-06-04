package main

import (
	"github.com/gin-gonic/gin"
	"github.com/konarkkapil1/gin-api/router"
	"github.com/konarkkapil1/gin-api/server"
)

func main() {
	gin := gin.Default()
	server := server.NewServer(gin, ":3000")

	router.InitRouter(gin)

	server.ListenAndServe()
}
