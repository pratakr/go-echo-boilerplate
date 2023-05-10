package routes

import (
	"kancha-api/app/caches"
	"kancha-api/app/database"
	"kancha-api/app/healthcheck"
	"kancha-api/app/http/controllers"
	"kancha-api/app/middlewares"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func DefineRouteAPI(e *echo.Echo, logger *zap.Logger) {

	db := database.ConnectMysql()
	redisClient := caches.Connect()
	dbResource := healthcheck.DbResource{DB: db}
	redisResource := healthcheck.RedisResource{Client: redisClient}
	controller := controllers.NewController(db, redisClient, logger)

	e.Use(middlewares.New(dbResource, redisResource))

	api := e.Group("/api")

	//auth
	api.POST("/login", controller.Login)
	api.GET("/me", controller.Profile)

	// v1
	v1 := api.Group("/v1")
	v1.GET("/user/profile", controller.Profile)
	v1.GET("/users", controller.FindUsersPaginate)
	v1.POST("/user", controller.CreateUser)
	v1.PUT("/user/:id", controller.UpdateUser)
	v1.DELETE("/user/:id", controller.DeleteUser)

	v1.GET("/houses", controller.GetHouses) // ดึงข้อมูลบ้านทั้งหมด
	v1.GET("/house/:id", controller.GetHousebyId)
}
