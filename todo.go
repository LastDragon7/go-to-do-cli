package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) ValidateIndex(idx int) error {
	if idx < 0 || idx >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(idx int) error {
	t := *todos

	if err := t.ValidateIndex(idx); err != nil {
		return err
	}

	*todos = append(t[:idx], t[idx+1:]...)

	return nil
}

func (todos *Todos) toggle(idx int) error {
	t := *todos

	if err := t.ValidateIndex(idx); err != nil {
		return err
	}

	isCompleted := t[idx].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[idx].CompletedAt = &completionTime
	}

	t[idx].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(idx int, title string) error {
	t := *todos

	if err := t.ValidateIndex(idx); err != nil {
		return err
	}

	t[idx].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)

	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completed_at := ""

		if t.Completed {
			completed = "✔"
			if t.CompletedAt != nil {
				completed_at = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completed_at)
	}

	table.Render()

}
