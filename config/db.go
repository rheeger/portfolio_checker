package config

import (
	"fmt"

	_ "github.com/lib/pq"
	"gopkg.in/mgo.v2"
)

// database
var DB *mgo.Database

// Portfolios Collection
var PortfoliosDB *mgo.Collection

// Users colleciton
var UsersDB *mgo.Collection

func init() {
	// get a mongo sessions
	//s, err := mgo.Dial("mongodb://bond:moneypenny007@localhost/bookstore")
	s, err := mgo.Dial("mongodb://localhost/portfolio_checker")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	DB = s.DB("portfolio_checker")
	PortfoliosDB = DB.C("portfolios")
	UsersDB = DB.C("users")

	fmt.Println("You connected to your porfolios_checker database.")
}
