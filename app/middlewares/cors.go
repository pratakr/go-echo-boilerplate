package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func CORSWithConfig() middleware.CORSConfig{
  return middleware.CORSConfig{
	  AllowOrigins: []string{"*"},
	  AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
  }
}
