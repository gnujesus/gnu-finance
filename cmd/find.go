/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/gnujesus/gnu-finance/config"
	"github.com/gnujesus/gnu-finance/internal/data"
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

		log.Println(company)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	findCmd.Flags().BoolP("detailed", "d", true, "See detailed information about a company. Default: true")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
