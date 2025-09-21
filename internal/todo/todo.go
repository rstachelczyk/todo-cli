package todo

import (
	"encoding/json"
	"github.com/dustin/go-humanize"
	"log"
	"os"
	"time"
)

type Todo struct {
	Text       string
	Priority   int
	CreatedAt  time.Time
	CompleteBy time.Time
	Completed  bool
}

type TodoView struct {
	Text       string
	Priority   string
	CompleteBy string
	Completed  string
}

func (t Todo) ToView() TodoView {
	return TodoView{
		Text:      t.Text,
		Priority:  humanizePriority(t.Priority),
		Completed: map[bool]string{true: "✅", false: "❌"}[t.Completed],
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

func SaveTodos(filename string, todos []Todo) error {
	jsonTodos, err := json.Marshal(todos)
	if err != nil {
		return err
	}
	writeToFile(filename, []byte(string(jsonTodos)))
	return nil
}

func writeToFile(filename string, data []byte) error {
	permissions := os.FileMode(0644)
	err := os.WriteFile(filename, data, permissions)
	if err != nil {
		log.Fatalf("Error saving todos: %v", err)
	}
	return nil
}

func GetTodos(filename string) ([]Todo, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var todos []Todo
	if err := json.Unmarshal(bytes, &todos); err != nil {
		return todos, err
	}

	return todos, nil
}
