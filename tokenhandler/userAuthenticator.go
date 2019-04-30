package tokenhandler

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iaronaraujo/RedCoins/models"
	"github.com/labstack/echo"
)

var jwtKey = []byte("my_secret_key")

//Claims represents login claims
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

//GenerateToken creates a token to represent an user session
func GenerateToken(c echo.Context, userMail string, userPW string) error {
	result := models.UserModel.Find("email=?", userMail)
	count, err := result.Count()
	if count != 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	var users []models.User
	result.All(&users)
	user := users[0]
	if user.Password != userPW {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Wrong Password",
		})
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, map[string]string{
		"token": tokenString,
	})
}
