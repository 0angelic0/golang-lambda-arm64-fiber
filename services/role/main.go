// main.go
package main

import (
	"context"
	"log"
	"os"

	"demo/pkg/role/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
)

var (
	fiberLambda *fiberadapter.FiberLambda
	app         *fiber.App
)

// init the Fiber Server
func init() {
	log.Printf("Fiber cold start")
	app = fiber.New()
}

// Handler will deal with Fiber working with Lambda
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContext(ctx, req)
}

// main
func main() {
	if isServerless() {
		http.RegisterHandlersWithBaseURL(app, &http.RoleHandler{}, "/role")
		fiberLambda = fiberadapter.New(app)
		// Make the handler available for Remote Procedure Call by AWS Lambda
		lambda.Start(Handler)
	} else {
		http.RegisterHandlers(app, &http.RoleHandler{})
		log.Fatal(app.Listen(":3000"))
	}
}

// isServerless check whether the program are running in serverless
func isServerless() bool {
	_, isOffline := os.LookupEnv("IS_OFFLINE")
	if isOffline {
		return false
	}
	return true
}
