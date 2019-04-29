package controllers

import (
	"net/http"

	"github.com/iaronaraujo/RedCoins/models"
	"github.com/labstack/echo"
)

//GetReportsByUserID gets the reports of an user by its id
func GetReportsByUserID(c echo.Context) error {
	userID := c.FormValue("userID")

	result := models.ReportModel.Find("user_id=?", userID)
	count, err := result.Count()
	if count < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"messagem": err.Error(),
		})
	}
	var reports []models.Report
	result.All(&reports)
	return c.JSON(http.StatusAccepted, reports)
}
