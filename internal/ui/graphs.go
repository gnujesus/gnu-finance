package ui

import (
	"github.com/gnujesus/gnu-finance/internal/data"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"os"
)

func GraphView(company data.CompanyInfo) error {
	kline := charts.NewKLine()

	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    company.Symbol + " Stock Price",
			Subtitle: "Daily OHLC",
		}),
		// Adding a Tooltip and DataZoom makes stock charts much more usable
		charts.WithTooltipOpts(opts.Tooltip{Show: opts.Bool(true), Trigger: "axis"}),
		charts.WithDataZoomOpts(opts.DataZoom{Type: "slider"}),
	)

	dates := make([]string, 0)
	klineData := make([]opts.KlineData, 0)

	for _, point := range company.PriceHistory {
		dates = append(dates, point.Date)
		// Correctly pass [Open, Close, Low, High]

		klineData = append(klineData, opts.KlineData{
			Value: []interface{}{point.Open, point.Close, point.Low, point.High},
		})
	}

	kline.SetXAxis(dates).
		AddSeries("Price", klineData)

	f, err := os.Create("stock_chart.html")
	if err != nil {
		return err
	}
	defer f.Close()

	return kline.Render(f)
}
