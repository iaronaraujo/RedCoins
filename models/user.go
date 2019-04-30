package models

import (
	"time"

	"github.com/iaronaraujo/RedCoins/lib"
)

// User represents an user of RedCoins
type User struct {
	ID        int64      `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
}

// UserModel receives the DataBase table
var UserModel = lib.Sess.Collection("users")
