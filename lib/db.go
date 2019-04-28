package lib

import (
	"log"

	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

var config = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "root",
	Password: "password",
	Database: "redcoins",
}

// Sess creates a database session
var Sess sqlbuilder.Database

func init() {
	var err error

	Sess, err = mysql.Open(config)
	if err != nil {
		log.Fatal(err.Error())
	}

}
