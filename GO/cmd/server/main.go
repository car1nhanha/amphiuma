package main

import (
	"github.com/car1nhanha/go-blog/internal/db"
	"github.com/car1nhanha/go-blog/internal/maria"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	db.Connect()

	router := gin.Default()
	maria.RegisterRoutes(router)

	router.Run(":8080")
}
