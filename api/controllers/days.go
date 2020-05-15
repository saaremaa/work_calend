package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/saaremaa/work_calend/models"
)

// CheckDay - обработчик по точке входа "/api/v1/check_day/:date"
// Формирует ответ по проверке - является ли день рабочим?
func (a *Controllers) CheckDay(c echo.Context) error {
	response := &models.MsgIsWorkDay{}
	isHoliday, err := a.Repository.DB.IsHoliday(c.Param("date"))
	if err != nil {
		response.Error = err.Error()
		return echo.NewHTTPError(http.StatusBadRequest, response)
	}
	response.Date = c.Param("date")
	response.WorkDay = !isHoliday
	return c.JSON(http.StatusOK, &response)
}
