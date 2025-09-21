package command

import (
	"fmt"
	"time"

	"github.com/rstachelczyk/todo-cli/internal/todo"
	"github.com/spf13/cobra"
)

var priority int
var completed bool

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

	addCmd.Flags().BoolVarP(
		&completed,
		"completed",
		"c",
		false,
		"Set as completed",
	)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Brief description of todo",
	Long:  `Adding a todo`,
	Args:  cobra.ExactArgs(1),
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

		fmt.Println("Added new todos")
	},
}

func buildTodo(description string) todo.Todo {
	todo := todo.Todo{
		Text:       description,
		CreatedAt:  time.Now(),
		Priority:   2,
		CompleteBy: time.Now(),
		Completed:  completed,
	}
	todo.SetPriority(priority)
	return todo
}
