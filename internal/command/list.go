package command

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/rstachelczyk/todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Display list of todos",
	Long:  `Display list of todos`,
	Run: func(cmd *cobra.Command, args []string) {
		todos, err := todo.GetTodos(dataFile)

		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("You currently have no todos. \nAdd one with the `add` command to get started or specify the correct config file")
			} else {
				fmt.Printf("An unexpected error occurred while reading file: %v\n", err)
			}
			return
		}

		if len(todos) == 0 {
			fmt.Println("You currently have no todos. \nAdd one with the `add` command to get started!")
			return
		} else {
			table := tablewriter.NewWriter(os.Stdout)
			table.Header([]string{"DESCRIPTION", "PRIORITY", "COMPLETE BY", "DONE"})
			table.Bulk(formatTodos(todos))
			table.Render()
		}
	},
}

func formatTodos(todos []todo.Todo) []todo.TodoView {
	formatted := []todo.TodoView{}
	for _, todo := range todos {
		formatted = append(formatted, todo.ToView())
	}
	return formatted
}
