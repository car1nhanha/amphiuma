package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/car1nhanha/amphiuma/internal/files"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ginLambda *ginadapter.GinLambda

func init() {
	godotenv.Load(".env")

	// Cria router normal
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/v1/:user/*path", files.GetFileHandler)
	router.GET("/v1/:user", files.ListFiles)

	// Adapta o Gin para Lambda
	ginLambda = ginadapter.New(router)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	log.Println("✅ Lambda inicializada com sucesso.")
	lambda.Start(Handler)
}
