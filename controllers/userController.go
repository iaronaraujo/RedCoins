package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/iaronaraujo/RedCoins/models"
	"github.com/iaronaraujo/RedCoins/utils"
	"github.com/labstack/echo"
)

// CreateUser creates a redcoins user
func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	bday, _ := strconv.Atoi(c.FormValue("bday"))
	bmonth, _ := strconv.Atoi(c.FormValue("bmonth"))
	byear, _ := strconv.Atoi(c.FormValue("byear"))

	birthDate := time.Date(byear, time.Month(bmonth), bday, 0, 0, 0, 0, time.UTC)
	birthDate.IsZero()

	var user models.User
	user.Name = name
	user.Email = email
	user.Password = password
	user.BirthDate = birthDate

	if name != "" && email != "" && password != "" {
		if _, err := models.UserModel.Insert(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"messagem": "User created successfully!",
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"mensagem": "Name, email or password were empty",
	})

}

// BuyBitCoins buy bitcoins
func BuyBitCoins(c echo.Context) error {
	return doBitCoinTransaction(c, models.BuyBitCoins)
}

//SellBitCoins sell bitcoins
func SellBitCoins(c echo.Context) error {
	return doBitCoinTransaction(c, models.SellBitCoins)
}

func doBitCoinTransaction(c echo.Context, transType models.TransactionType) error {
	userID, _ := strconv.Atoi(c.FormValue("user_id"))
	count, _ := models.UserModel.Find("id=?", userID).Count()
	if count < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"mensagem": "User not found",
		})
	}

	currencyTyp := utils.CurrencyType(c.FormValue("currency"))
	bitcoins, _ := strconv.ParseFloat(c.FormValue("bitcoins"), 32)
	value := utils.ConvertBitcoinsToCurrency(float32(bitcoins), currencyTyp)
	day, _ := strconv.Atoi(c.FormValue("day"))
	month, _ := strconv.Atoi(c.FormValue("month"))
	year, _ := strconv.Atoi(c.FormValue("year"))
	transactionDate := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return createReport(c, transType, float32(bitcoins), value, currencyTyp, transactionDate, userID)
}

//CreateReport creates a transaction report
func createReport(c echo.Context, transType models.TransactionType, bitcoins float32, value float32, currTyp utils.CurrencyType, date time.Time, userID int) error {
	var report models.Report
	report.Transaction = transType
	report.BitCoins = bitcoins
	report.Value = value
	report.Currency = currTyp
	report.TransactionDate = date
	report.UserID = userID

	if _, err := models.ReportModel.Insert(report); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, report)

}
