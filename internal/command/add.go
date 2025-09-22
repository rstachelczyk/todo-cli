package command

import (
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/rstachelczyk/todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

var priority int
var done bool
var due int

func init() {
	rootCmd.AddCommand(addCmd)
	// TODO: Add string priority flag too (high, medium, low vs 1, 2, 3)
	addCmd.Flags().IntVarP(
		&priority,
		"priority",
		"p",
		2,
		"Priority: 1 (High), 2 (Medium), 3 (Low)",
	)

	addCmd.Flags().BoolVar(
		&done,
		"done",
		false,
		"Set as done",
	)

	addCmd.Flags().IntVar(
		&due,
		"due",
		0,
		`Set due date in number of days. (default 0)
		...
         -1 = yesterday 
          0 = today 
          1 = tomorrow
		...
		`,
	)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Brief description of todo",
	Long:  `Adding a todo`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		newTodos := []todo.Todo{}
		for _, description := range args {
			newTodo := buildTodo(description)
			newTodos = append(newTodos, newTodo)
		}
		err := todo.SaveTodos(dataFile, newTodos)
		if err != nil {
			fmt.Println(fmt.Errorf("%v", err))
		}

		fmt.Println("Added new todos:")
		table := tablewriter.NewWriter(os.Stdout)
		table.Header([]string{"DESCRIPTION", "PRIORITY", "COMPLETE BY", "DONE"})
		table.Bulk(formatTodos(newTodos))
		table.Render()
	},
}

func buildTodo(description string) todo.Todo {
	todo := todo.Todo{
		Text:       description,
		CreatedAt:  time.Now(),
		Priority:   2,
		CompleteBy: EndOfDay().AddDate(0, 0, due),
		Done:       done,
	}
	todo.SetPriority(priority)
	return todo
}

func EndOfDay() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
}
