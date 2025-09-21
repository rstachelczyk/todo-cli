package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var dataFile string

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a CLI tool for managing your to-do list",
	Long:  `A fast and simple CLI tool to manage your to-do list. It will help you stay organized and get more done in less time. It's designed to be as simple as possible to help you stay on top of things.`,
}

func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Unable to detect home directory. Please set datafile using --datafile.")
	}

	rootCmd.PersistentFlags().StringVar(
		&dataFile,
		"datafile",
		home+string(os.PathSeparator)+"todos.json",
		"data file to store todos",
	)
}

// func Execute(key string) {
func Execute() {

	// if dataFile == "" {
	// 	fmt.Println("CURRENCY_API_KEY environment variable is not set")
	// 	os.Exit(1)
	// }
	// apiKey = key
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
