package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/iaronaraujo/RedCoins/models"
	"github.com/iaronaraujo/RedCoins/tokenhandler"
	"github.com/iaronaraujo/RedCoins/utils"
	"github.com/labstack/echo"
)

// CreateUser creates a RedCoins user
func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	userType := models.UserType(c.FormValue("type"))
	bday, _ := strconv.Atoi(c.FormValue("bday"))
	bmonth, _ := strconv.Atoi(c.FormValue("bmonth"))
	byear, _ := strconv.Atoi(c.FormValue("byear"))

	birthDate := time.Date(byear, time.Month(bmonth), bday, 0, 0, 0, 0, time.UTC)
	birthDate.IsZero()

	var user models.User
	var err error
	user.Name = name
	user.Email = email
	user.Password, err = utils.HashPassword(password)
	user.Type = userType
	user.BirthDate = birthDate

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if name != "" && email != "" && password != "" {
		if _, err := models.UserModel.Insert(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]string{
			"message": "User created successfully!",
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"message": "Name, email or password were empty",
	})

}

//Login signs an user in, returning an access token
func Login(c echo.Context) error {
	userMail := c.FormValue("email")
	userPW := c.FormValue("password")

	result := models.UserModel.Find("email=?", userMail)
	count, err := result.Count()
	if count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "There isnt an user with this email",
		})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	var users []models.User
	result.All(&users)
	user := users[0]
	if !utils.CheckPasswordHash(userPW, user.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Wrong Password",
		})
	}

	tokenString := tokenhandler.GenerateToken(user)
	return c.JSON(http.StatusAccepted, map[string]string{
		"token": tokenString,
	})

}

// BuyBitCoins is the operation an user can make to buy BitCoins, which generates a report
func BuyBitCoins(c echo.Context) error {
	return doBitCoinTransaction(c, models.BuyBitCoins)
}

//SellBitCoins is the operation an user can make to sell BitCoins, which generates a report.
func SellBitCoins(c echo.Context) error {
	return doBitCoinTransaction(c, models.SellBitCoins)
}

//doBitCoinTransaction executes a buy or sell BitCoins transaction
func doBitCoinTransaction(c echo.Context, transType models.TransactionType) error {
	token := c.Request().Header.Get("token")
	userID, _ := tokenhandler.GetLoggedUser(token)
	count, _ := models.UserModel.Find("id=?", userID).Count()
	if count < 1 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "User not found",
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

//createReport creates a transaction report
func createReport(c echo.Context, transType models.TransactionType, bitcoins float32, value float32, currTyp utils.CurrencyType, date time.Time, userID int64) error {
	var report models.Report
	report.Transaction = transType
	report.BitCoins = bitcoins
	report.Value = value
	report.Currency = currTyp
	report.TransactionDate = date
	report.UserID = userID

	newID, err := models.ReportModel.Insert(report)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	id, _ := newID.(int64)
	report.ID = id

	return c.JSON(http.StatusCreated, report)

}
