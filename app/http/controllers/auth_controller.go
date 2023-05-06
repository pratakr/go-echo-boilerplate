package controllers

import (
	"errors"
	"kancha-api/app/http/request"
	"kancha-api/app/http/response"
	"kancha-api/app/model"
	"kancha-api/app/service"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	ID     int64  `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	RoleID int32  `json:"role_id"`
	jwt.RegisteredClaims
}

func loginFailed(c echo.Context) error {
	errResponse := &response.ErrorResponse{
		Error: "Failed Login",
	}
	return c.JSON(http.StatusForbidden, errResponse)
}

func (con *controller) Login(c echo.Context) error {
	req := &request.AuthLoginRequest{}
	if err := c.Bind(&req); err != nil {
		con.Logger.Warn("bind data failed!")
		return loginFailed(c)
	}

	var user *model.User
	newService := service.NewService(con.Db)

	user, err := newService.UserLogin(req.Email, req.Password)
	if err != nil {
		con.Logger.Warn("login failed!")
		return loginFailed(c)
	}

	token, err := con.CreateToken(*user)

	userResponse := &response.UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RoleID: user.RoleID,
		Token:  token,
	}

	if err != nil {
		con.Logger.Warn("login failed!")
		return loginFailed(c)
	}

	return c.JSON(http.StatusOK, userResponse)
}

func (con *controller) Profile(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return errors.New("JWT token missing or invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("JWT token missing or invalid")
	}

	return c.JSON(http.StatusOK, claims)
}

func (con *controller) CreateToken(user model.User) (string, error) {

	key := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := &jwtCustomClaims{
		user.ID,
		user.Email,
		user.Name,
		user.RoleID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(key)
	return t, err
}
