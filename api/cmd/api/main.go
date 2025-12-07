package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/car1nhanha/amphiuma/internal/files"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ginLambda *ginadapter.GinLambdaV2
var router *gin.Engine

func init() {
	godotenv.Load(".env")

	// Cria router normal
	router = gin.Default()
	router.Use(cors.Default())

	router.GET("/:user/:repo/*path", files.GetFileHandler)
	router.GET("/:user", files.ListFiles)
	router.GET("/", (func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "✅👍😃👍",
		})
	}))

	// Adapta o Gin para Lambda
	ginLambda = ginadapter.NewV2(router)
}

func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Verifica se está rodando localmente ou na Lambda
	if os.Getenv("IS_RUNNING_LOCAL") == "true" {
		// Modo local - inicia servidor HTTP normal
		log.Println("🚀 Rodando localmente em http://localhost:8080")
		if err := router.Run(":8080"); err != nil {
			log.Fatal("Erro ao iniciar servidor:", err)
		}
	} else {
		// Modo Lambda
		log.Println("✅ Lambda inicializada com sucesso.")
		lambda.Start(Handler)
	}
}
