package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/saaremaa/work_calend/models"
)

// Statistic cтруктура для статистических данных
type Statistic struct {
	MinYear    string
	MaxYear    string
	LastUpdate string
}

// RespROOT обработчик по точке входа "/"
func (a *Controllers) RespROOT(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "https://enicom.ru/")
}

// RespPing обработчик по точке входа "/api/v1/ping"
func (a *Controllers) Health(c echo.Context) error {
	response := &models.MessageUnauthorized{
		Error:   "",
		Message: "ok",
	}
	return c.JSON(http.StatusOK, &response)
}

// Stat обработчик по точке входа "/api/v1/stat"
func (a *Controllers) Stat(c echo.Context) error {
	stat := a.Repository.DB.CaclStat()
	response := &models.MsgStat{
		Error:      "",
		StartYear:  stat.MinYear,
		StopYear:   stat.MaxYear,
		LastUpdate: a.StartTime.Format("02/01/2006 15:04:05"),
		UPTime:     time.Since(a.StartTime).String(),
	}
	return c.JSON(http.StatusOK, &response)
}
