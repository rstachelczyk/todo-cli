package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// var apiKey string

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a CLI tool for managing your to-do list",
	Long:  `A fast and simple CLI tool to manage your to-do list`,
}

// func Execute(key string) {
func Execute() {

	// if key == "" {
	// 	fmt.Println("CURRENCY_API_KEY environment variable is not set")
	// 	os.Exit(1)
	// }
	// apiKey = key
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
