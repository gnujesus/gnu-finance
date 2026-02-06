/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/gnujesus/gnu-finance/config"
	"github.com/gnujesus/gnu-finance/internal/data"
	"github.com/gnujesus/gnu-finance/internal/tui"
	"github.com/gnujesus/gnu-finance/internal/ui"
	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find a company",
	Long:  `After providing the stock tick symbol, finds a the company details using verified APIs.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Init()
		alphaVantage := data.AlphaVantage{ApiKey: cfg.ApiKey}

		query := data.Query{
			CompanySymbol: args[0],
			DateRange:     "daily",
		}

		company, err := alphaVantage.Fetch(query)

		if err != nil {
			log.Fatal("Error: ", err)
			return
		}

		model := tui.InitialModel(company)

		p := tea.NewProgram(model)

		finalModel, err := p.Run()

		if err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

		m := finalModel.(tui.Model)

		switch m.Cursor {
		case 0:
			tui.SimpleView(m.Company)

		case 1:
			tui.DetailedView(m.Company)

		case 2:
			// TODO: CURRENTLY NOT WORKING
			ui.GraphView(m.Company)
			exec.Command("open", "stock_charts.html").Start()

		case 3:
			os.Exit(0)

		default:
			fmt.Println("Default behavior")
		}

		if _, err := p.Run(); err != nil {
			if _, err := p.Run(); err != nil {
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	findCmd.Flags().BoolP("detailed", "d", false, "See detailed information about a company. Default: true")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
