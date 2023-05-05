package main

import (
	"log"

	"kancha-api/app/middlewares"
	"kancha-api/app/routes"

	"github.com/joho/godotenv"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

// var ctx = context.Background()

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}
	// Echo instance
	e := echo.New()

	logger, _ := zap.NewProduction()

	// Middleware
	e.Use(middlewares.ZapLogger(logger))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middlewares.CORSWithConfig()))
	e.Use(echojwt.WithConfig(middlewares.JwtConfig()))

	routes.DefineRouteAPI(e, logger)

	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
