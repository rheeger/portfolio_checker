package portfolio

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/rheeger/portfolio_checker/config"
	"gopkg.in/mgo.v2/bson"
)

type Holding struct {
	Name   string
	Ticker string
	Amount float64
}

func AllHoldings() ([]Holding, error) {
	hldgs := []Holding{}
	err := config.PortfoliosDB.Find(bson.M{}).All(&hldgs)
	if err != nil {
		return nil, err
	}
	return hldgs, nil
}

func OneHolding(r *http.Request) (Holding, error) {
	hldgs := Holding{}
	Ticker := r.FormValue("Ticker")
	if Ticker == "" {
		return hldgs, errors.New("400. Bad Request.")
	}
	err := config.PortfoliosDB.Find(bson.M{"Ticker": Ticker}).One(&hldgs)
	if err != nil {
		return hldgs, err
	}
	return hldgs, nil
}

func PutHolding(r *http.Request) (Holding, error) {
	// get form values
	hldgs := Holding{}
	hldgs.Ticker = r.FormValue("Ticker")
	hldgs.Name = r.FormValue("Name")
	amnt := r.FormValue("Amount")

	// validate form values
	if hldgs.Ticker == "" || hldgs.Name == "" || amnt == "" {
		return hldgs, errors.New("400. Bad request. All fields must be complete.")
	}
	// convert amount form value to float 64
	f64, err := strconv.ParseFloat(amnt, 64)
	if err != nil {
		return hldgs, errors.New("406. Not Acceptable. Price must be a number.")
	}
	hldgs.Amount = float64(f64)

	// insert values
	err = config.PortfoliosDB.Insert(hldgs)
	if err != nil {
		return hldgs, errors.New("500. Internal Server Error." + err.Error())
	}
	return hldgs, nil
}

func UpdateHolding(r *http.Request) (Holding, error) {
	// get form values
	hldgs := Holding{}
	hldgs.Ticker = r.FormValue("Ticker")
	hldgs.Name = r.FormValue("Name")
	amnt := r.FormValue("Amount")

	if hldgs.Ticker == "" || hldgs.Name == "" || amnt == "" {
		return hldgs, errors.New("400. Bad Request. Fields can't be empty.")
	}

	// convert amount form value to float 64
	f64, err := strconv.ParseFloat(amnt, 64)
	if err != nil {
		return hldgs, errors.New("406. Not Acceptable. Price must be a number.")
	}
	hldgs.Amount = float64(f64)

	// update values
	err = config.PortfoliosDB.Update(bson.M{"Ticker": hldgs.Ticker}, &hldgs)
	if err != nil {
		return hldgs, err
	}
	return hldgs, nil
}

func DeleteHolding(r *http.Request) error {
	Ticker := r.FormValue("Ticker")
	if Ticker == "" {
		return errors.New("400. Bad Request.")
	}

	err := config.PortfoliosDB.Remove(bson.M{"Ticker": Ticker})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
