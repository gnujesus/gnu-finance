package tui

import (
	"fmt"
	"github.com/gnujesus/gnu-finance/internal/data"
)

func SimpleView(company data.CompanyInfo) {
	fmt.Printf("\n%s (%s)\n", company.Name, company.Symbol)
	fmt.Printf("Last Updated: %s %s\n\n", company.LastRefreshed, company.TimeZone)

	if len(company.PriceHistory) > 0 {
		latest := company.PriceHistory[0]
		fmt.Printf("Latest Price: $%.2f\n", latest.Close)
		fmt.Printf("Open: $%.2f | High: $%.2f | Low: $%.2f\n", latest.Open, latest.High, latest.Low)
		fmt.Printf("Volume: %d\n\n", latest.Volume)
	}
}
