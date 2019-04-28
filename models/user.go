package models

import (
	"time"
)

// User represents an user of RedCoins
type User struct {
	ID        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"password"`
	BirthDate time.Time `db:"birth_date" json:"birth_date"`
}
