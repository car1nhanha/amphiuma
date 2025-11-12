package main

import (
	"context"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/car1nhanha/amphiuma/internal/files"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ginLambda *ginadapter.GinLambdaV2

func init() {
	godotenv.Load(".env")

	// Cria router normal
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/:user/*path", files.GetFileHandler)
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
	log.Println("✅ Lambda inicializada com sucesso.")
	lambda.Start(Handler)
}
