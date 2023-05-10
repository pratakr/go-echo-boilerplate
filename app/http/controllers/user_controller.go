package controllers

import (
	"fmt"
	"kancha-api/app/http/request"
	"kancha-api/app/http/response"
	"kancha-api/app/model"
	"kancha-api/app/service"
	"kancha-api/app/utils"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"gopkg.in/validator.v2"
)

func (con *controller) CreateUser(c echo.Context) error {
	var req request.CreateUserRequest
	err := c.Bind(&req)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	err = validator.Validate(req)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	newService := service.NewService(con.Db)
	var userModel model.User
	copier.Copy(&userModel, &req)
	userModel.EmailVerifiedAt = nil
	con.Logger.Debug(fmt.Sprintf("userModel: %v", userModel))

	if newService.ExistUser(&userModel) {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "user already exist"})
	}

	user, err := newService.CreateUser(&userModel)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var res response.UserResponse
	copier.Copy(&res, &user)
	return c.JSON(http.StatusOK, res)
}

func (con *controller) UpdateUser(c echo.Context) error {
	var req request.UpdateUserRequest
	err := c.Bind(&req)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	userID := c.Param("id")

	err = validator.Validate(req)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	newService := service.NewService(con.Db)
	var userModel model.User
	copier.Copy(&userModel, &req)
	userModel.ID, err = strconv.ParseInt(userID, 10, 64)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}
	userModel.EmailVerifiedAt = nil
	con.Logger.Debug(fmt.Sprintf("userModel: %v", userModel))

	user, err := newService.UpdateUser(&userModel)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var res response.UserResponse
	copier.Copy(&res, &user)
	return c.JSON(http.StatusOK, res)
}

func (con *controller) DeleteUser(c echo.Context) error {
	id := c.Param("id")
	con.Logger.Debug(fmt.Sprintf("id: %v", id))
	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	newService := service.NewService(con.Db)
	err = newService.DeleteUser(userID)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response.ErrorResponse{Message: "success"})
}

func (con *controller) FindUsersPaginate(c echo.Context) error {
	pagination := utils.NewPagination(c)

	newService := service.NewService(con.Db)

	pagination, err := newService.FindUsersPaginate(pagination)
	if err != nil {
		con.Logger.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, pagination)
}
