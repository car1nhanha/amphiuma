package main

import (
	"github.com/car1nhanha/amphiuma/internal/files"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(cors.Default())

	{
		v1 := router.Group("/v1")
		v1.GET("/:user/:path", files.GetFileHandler)
	}

	router.Run(":8080")
}
