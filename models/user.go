package models

import (
	"time"

	"github.com/iaronaraujo/RedCoins/lib"
)

//UserType is the type of an user (ADMIN or NORMAL)
type UserType string

const (
	//ADMIN represents an administrator user of RedCoins, which is capable of seeing the reports
	ADMIN UserType = "ADMIN"
	//NORMAL represents a normal user of RedCoins, which can only buy and sell bitcoins
	NORMAL UserType = "NORMAL"
)

// User represents an user of RedCoins
type User struct {
	ID        int64     `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	Type      UserType  `db:"type" json:"type"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
}

// UserModel receives the DataBase table
var UserModel = lib.Sess.Collection("users")
