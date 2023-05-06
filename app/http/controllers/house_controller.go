package controllers

import (
	"errors"
	"kancha-api/app/http/response"
	"kancha-api/app/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (con *controller) GetHouses(c echo.Context) error {
	newService := service.NewService(con.Db)
	houses, err := newService.FindHouses()
	if err != nil {
		con.Logger.Warn(err.Error())
		return errors.New("find houses failed")
	}

	var houseResponses []*response.HouseResponse
	for _, house := range houses {
		houseResponses = append(houseResponses, &response.HouseResponse{
			ID:          house.ID,
			Code:        house.Code,
			Name:        house.Name,
			Temperature: house.Temperature,
			Humidity:    house.Humidity,
			Light:       house.Light,
			UpdatedAt:   house.UpdatedAt,
		})
	}
	return c.JSON(http.StatusOK, houseResponses)
}

func (con *controller) GetHousebyId(c echo.Context) error {
	newService := service.NewService(con.Db)
	house, err := newService.FindHouseById(c.Param("id"))
	if err != nil {
		con.Logger.Warn(err.Error())
		return errors.New("find house failed")
	}

	houseResponse := &response.HouseResponse{
		ID:          house.ID,
		Code:        house.Code,
		Name:        house.Name,
		Temperature: house.Temperature,
		Humidity:    house.Humidity,
		Light:       house.Light,
		UpdatedAt:   house.UpdatedAt,
	}
	return c.JSON(http.StatusOK, houseResponse)
}
