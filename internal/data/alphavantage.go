package data

import (
	"encoding/json"
	"fmt"
	"log"
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

	// TODO: REMEMBER TO PARSE THE TIMESERIES SINCE IT IS A STRING IN THE RES, BUT 
	FLOAT IN THE MODEL
	 
	return CompanyInfo{
		Name:          "Filler",
		Symbol:        raw.MetaData.Symbol,
		LastRefreshed: raw.MetaData.LastRefreshed,
		TimeZone:      raw.MetaData.TimeZone,
		PriceHistory:  raw.TimeSeries,
	}, nil

}
