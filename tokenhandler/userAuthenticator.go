package tokenhandler

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/iaronaraujo/RedCoins/models"
)

var jwtKey = []byte("my_secret_key")

//Claims represents login claims
type Claims struct {
	UserID   int64           `json:"user_id"`
	UserType models.UserType `json:"user_type"`
	jwt.StandardClaims
}

//GenerateToken creates a token to represent an user session
func GenerateToken(user models.User) string {

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID:   user.ID,
		UserType: user.Type,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)
	return tokenString
}

//GetLoggedUser returns the id of the user related to the token or -1 if the token isn't valid
func GetLoggedUser(token string) (int64, models.UserType) {

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid || err != nil {
		return -1, models.NORMAL
	}
	return claims.UserID, claims.UserType

}
