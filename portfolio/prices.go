package portfolio

import (
	"os"

	cmc "github.com/miguelmota/go-coinmarketcap/pro/v1"
)

func GetPrice(hldg Holding) float64 {
	client := cmc.NewClient(&cmc.Config{
		ProAPIKey: os.Getenv("CMC_PRO_API_KEY"),
	})

	quotes, err := client.Cryptocurrency.LatestQuotes(&cmc.QuoteOptions{
		Symbol:  hldg.Ticker,
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
