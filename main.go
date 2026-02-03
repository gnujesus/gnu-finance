/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/gnujesus/gnu-finance/cmd"
	"github.com/gnujesus/gnu-finance/config"
)

func main() {
	cfg := config.Init()

	cmd.Execute()
}
