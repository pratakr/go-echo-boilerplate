package controllers

import (
	"kancha-api/app/http/request"
	"kancha-api/app/http/response"
	"kancha-api/app/model"
	"kancha-api/app/service"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

func (con *controller) CreateUser(c echo.Context) error {
	var req request.CreateUserRequest
	err := c.Bind(&req)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.String(http.StatusBadRequest, "bad request")
	}

	err = c.Validate(req)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.String(http.StatusBadRequest, err.Error())
	}

	newService := service.NewService(con.Db)
	var userModel model.User
	copier.Copy(&userModel, &req)
	user, err := newService.CreateUser(&userModel)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.String(http.StatusInternalServerError, "create user failed")
	}

	var res response.UserResponse
	copier.Copy(&res, &user)
	return c.JSON(http.StatusOK, res)
}
