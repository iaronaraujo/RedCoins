package controllers

import (
    "time"
    "strconv"
    "net/http"
    "github.com/labstack/echo"
    "github.com/iaronaraujo/RedCoins/models"
)

func CreateUser(c echo.Context) error {
    name := c.FormValue("name")
    email := c.FormValue("email")
    password := c.FormValue("password")
    bday := c.FormValue("bday")
    bmonth := c.FormValue("bmonth")
    byear := c.FormValue("byear")

    birthDate := time.Date(strconv.Atoi(byear),
    strconv.Atoi(bmonth),
    strconv.Atoi(bday),
    0, 0, 0, time.UTC)

    var user models.User
    user.Name = name
    user.Email = email
    user.Password = password
    user.BirthDate = birthDate

    if nome != "" && email != "" && password != "" {
        if _, err := models.UserModel.Insert(user); err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Creation of a new user was not possible",
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