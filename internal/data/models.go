package data

type Query struct {
	CompanySymbol string // company stock tick symbol
	DateRange     string // date range to check for the company. Simulating google results, in which you check for today, last 3 days, last month, last year or all time
}

type PricePoint struct {
	Date   string
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume int64
}

type CompanyInfo struct {
	Name          string
	Symbol        string
	LastRefreshed string
	TimeZone      string
	PriceHistory  []PricePoint
}
