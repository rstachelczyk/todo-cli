package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/dustin/go-humanize"
)

type Todo struct {
	Text       string    `json:"description"`
	Priority   int       `json:"priority"`
	CreatedAt  time.Time `json:"created_at"`
	CompleteBy time.Time `json:"complete_by"`
	Done       bool      `json:"done"`
}

type TodoView struct {
	Text       string
	Priority   string
	CompleteBy string
	Done       string
}

func (t Todo) ToView() TodoView {
	return TodoView{
		Text:     t.Text,
		Priority: humanizePriority(t.Priority),
		Done:     map[bool]string{true: "X", false: ""}[t.Done],
		// CompleteBy: humanizeDate(t.CompleteBy),
		CompleteBy: humanize.Time(t.CompleteBy),
	}
}

// helper for priority formatting
func humanizePriority(p int) string {
	switch p {
	case 1:
		return "High"
	case 3:
		return "Low"
	default:
		return "Medium"
	}
}

// helper for date formatting
func humanizeDate(d time.Time) string {
	now := time.Now()
	today := now.Truncate(24 * time.Hour)
	completeDay := d.Truncate(24 * time.Hour)

	switch completeDay.Sub(today) {
	case 0:
		return "Today"
	case 24 * time.Hour:
		return "Tomorrow"
	case -24 * time.Hour:
		return "Yesterday"
	default:
		// TODO:Add local support?
		return d.Format(time.DateOnly)
	}
}

func (todo *Todo) SetPriority(priority int) {
	switch priority {
	case 1:
		todo.Priority = 1
	case 3:
		todo.Priority = 3
	default:
		todo.Priority = 2
	}
}

// func SaveTodos(filename string, todos []Todo) error {
// 	jsonTodos, err := json.Marshal(todos)
// 	if err != nil {
// 		return err
// 	}
// 	writeToFile(filename, []byte(string(jsonTodos)))
// 	return nil
// }

// func writeToFile(filename string, data []byte) error {
// 	permissions := os.FileMode(0644)
// 	err := os.WriteFile(filename, data, permissions)
// 	if err != nil {
// 		log.Fatalf("Error saving todos: %v", err)
// 	}
// 	return nil
// }

func SaveTodos(filename string, newTodos []Todo) error {
	var todos []Todo

	// Read existing todos
	data, err := os.ReadFile(filename)
	if err == nil {
		if len(data) > 0 {
			if err := json.Unmarshal(data, &todos); err != nil {
				return fmt.Errorf("failed to unmarshal existing todos: %w", err)
			}
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("failed to read file: %w", err)
	}

	// Append new todos
	todos = append(todos, newTodos...)

	// Marshal and write back
	jsonTodos, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonTodos, 0644)
}

func GetTodos(filename string) ([]Todo, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var todos []Todo
	if err := json.Unmarshal(bytes, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}
