package ui

import (
	"os"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/gnujesus/gnu-finance/internal/data"
)

func GraphView(company data.CompanyInfo) error {
	// Create a new line chart
	line := charts.NewLine()
	
	// Set global options
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    company.Symbol + " Stock Price",
			Subtitle: "Close prices over time",
		}),
	)
	
	// Prepare data
	dates := make([]string, 0)
	prices := make([]opts.LineData, 0)
	
	for _, point := range company.PriceHistory {

		dates = append(dates, point.Date)
		prices = append(prices, opts.LineData{Value: point.Close})
	}
	
	// Add data to chart
	line.SetXAxis(dates).
		AddSeries("Close Price", prices).

		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}))
	

	// Render to HTML file
	f, err := os.Create("stock_chart.html")
	if err != nil {
		return err
	}
	defer f.Close()
	
	return line.Render(f)
}
