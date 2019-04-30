package controllers

import (
	"net/http"

	"strconv"
	"time"

	"github.com/iaronaraujo/RedCoins/models"
	"github.com/iaronaraujo/RedCoins/tokenhandler"
	"github.com/labstack/echo"
)

//GetReportsByUserID gets the reports of an user by its id
func GetReportsByUserID(c echo.Context) error {
	token := c.Request().Header.Get("token")
	_, userType := tokenhandler.GetLoggedUser(token)

	if userType != models.ADMIN {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You need to be an admin to view the reports",
		})
	}

	userID := c.FormValue("userID")
	result := models.ReportModel.Find("user_id=?", userID)
	count, _ := result.Count()
	if count < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "User not found",
		})
	}
	var reports []models.Report
	result.All(&reports)
	return c.JSON(http.StatusAccepted, reports)
}

//GetReportsByDate gets the reports of an user by its id
func GetReportsByDate(c echo.Context) error {
	token := c.Request().Header.Get("token")
	_, userType := tokenhandler.GetLoggedUser(token)

	if userType != models.ADMIN {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "You need to be an admin to view the reports",
		})
	}

	day, _ := strconv.Atoi(c.FormValue("day"))
	month, _ := strconv.Atoi(c.FormValue("month"))
	year, _ := strconv.Atoi(c.FormValue("year"))
	reportDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	result := models.ReportModel.Find("transaction_date=?", reportDate)
	count, err := result.Count()
	if count < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	var reports []models.Report
	result.All(&reports)
	return c.JSON(http.StatusAccepted, reports)
}
