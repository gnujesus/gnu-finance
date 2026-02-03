package data

import (
	"fmt"
	"log"
	"strconv"

	"encoding/json"
	"net/http"
)

// these mirror the raw API response exactly.
//
//	different from what I have on the models.
//
// the models are the application models, clean and neat
// this ones mirror the api response, so they can be parse and put into the
// application models.
type alphaResponse struct {
	MetaData   alphaMetaData              `json:"Meta Data"`
	TimeSeries map[string]alphaPricePoint `json:"Time Series (Daily)"`
}

type alphaMetaData struct {
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	TimeZone      string `json:"5. Time Zone"`
}

type alphaPricePoint struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

type AlphaVantage struct {
	ApiKey string
}

func (a AlphaVantage) Fetch(q Query) (CompanyInfo, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s", q.CompanySymbol, a.ApiKey)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Error fetching company information from API.")
		return CompanyInfo{}, err
	}

	var raw alphaResponse
	err = json.NewDecoder(res.Body).Decode(&raw)
	if err != nil {
		log.Fatal("Error parsing company information into JSON.")
		return CompanyInfo{}, err
	}

	history := []PricePoint{}

	for date, point := range raw.TimeSeries {
		// using openP because close is a reserved keyword, so I use P on open just to make it a convention between
		// open and close. This will only be used here and when a loop is needed. Don't mind it too much.
		openP, _ := strconv.ParseFloat(point.Open, 64)
		high, _ := strconv.ParseFloat(point.High, 64)
		low, _ := strconv.ParseFloat(point.Low, 64)
		closeP, _ := strconv.ParseFloat(point.Close, 64)
		volume, _ := strconv.ParseInt(point.Volume, 10, 64)

		history = append(history, PricePoint{
			Date:   date,
			Open:   openP,
			High:   high,
			Low:    low,
			Close:  closeP,
			Volume: volume,
		})
	}

	return CompanyInfo{
		Name:          "Filler",
		Symbol:        raw.MetaData.Symbol,
		LastRefreshed: raw.MetaData.LastRefreshed,
		TimeZone:      raw.MetaData.TimeZone,
		PriceHistory:  history,
	}, nil

}
