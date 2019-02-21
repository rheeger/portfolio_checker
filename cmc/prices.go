package cmc

import (
	"os"

	"github.com/rheeger/portfolio_checker/portfolio"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

func GetPrice(hldg portfolio.Holding) float64 {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: os.Getenv("CMC_PRO_API_KEY"),
	})

	quotes, err := client.Cryptocurrency.LatestQuotes(&cmc.QuoteOptions{
		Symbol:  hldg.BinanceTicker,
		Convert: "USD",
	})
	if err != nil {
		panic(err)
	}

	price := 0.00

	for _, quote := range quotes {
		price = quote.Quote["USD"].Price
	}

	return price
}
