package controllers

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (con *controller) ClearCache(c echo.Context) error {
	//cacheKey := fmt.Sprintf("HomeBanner")
	result,err := con.RedisClient.FlushAll(context.Background()).Result()
	if err != nil {
		con.Logger.Error(err.Error())
	}
	return c.JSON(http.StatusOK,result)
}
