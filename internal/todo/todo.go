package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// TODO: Add CompletedAt when I add `done` command
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
		Text:       t.Text,
		Priority:   humanizePriority(t.Priority),
		Done:       map[bool]string{true: "X", false: ""}[t.Done],
		CompleteBy: humanizeDate(t.CompleteBy),
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
func humanizeDate(t time.Time) string {
	now := time.Now()
	// To perform a day-only comparison, we must strip the time components.
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	target := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	// Calculate the difference in days.
	diffDays := int(target.Sub(today).Hours() / 24)

	switch {
	case diffDays == 0:
		return "Today"
	case diffDays == 1:
		return "Tomorrow"
	case diffDays == -1:
		return "Yesterday"
	case diffDays > 1 && diffDays <= 7:
		return fmt.Sprintf("%d days from now", diffDays)
	case diffDays < -1 && diffDays >= -7:
		return fmt.Sprintf("%d days ago", -diffDays)
	default:
		// For dates outside the 7-day range, fall back to a standard format.
		return t.Format("1/2/06")
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
