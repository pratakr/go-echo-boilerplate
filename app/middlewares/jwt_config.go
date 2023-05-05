package middlewares

import (
	"kancha-api/app/utils"
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func JwtConfig() echojwt.Config {

	signingKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	return echojwt.Config{
		SigningKey: signingKey,
		Skipper: func(c echo.Context) bool {
			// Skip middleware if path is equal 'login'
			arr := []string{
				"/api/login",
				"/api/register",
				"/healthz",
			}
			path := c.Request().URL.Path
			if utils.Contains(arr, path) {
				return true
			}
			return false
		},
	}
}
