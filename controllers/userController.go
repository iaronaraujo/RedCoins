package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/iaronaraujo/RedCoins/models"
	"github.com/labstack/echo"
)

// CreateUser creates a redcoins user
func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")
	bday := c.FormValue("bday")
	bmonth := c.FormValue("bmonth")
	byear := c.FormValue("byear")

	yearNum, _ := strconv.Atoi(byear)
	monthNum, _ := strconv.Atoi(bmonth)
	dayNum, _ := strconv.Atoi(bday)

	birthDate := time.Date(yearNum, time.Month(monthNum), dayNum, 0, 0, 0, 0, time.UTC)
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
				//"message": "Creation of a new user was not possible",
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
