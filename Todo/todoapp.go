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

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:     title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index <0 || index >= len(*todos) {
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	Todos := *todos

	if err := Todos.validateIndex(index); err != nil {
		return err
	}

	*todos = append(Todos[:index], Todos[index + 1:]...)

	return nil
}

func (todos *Todos) Toggle(index int) error {
	Todos := *todos
	if err := Todos.validateIndex(index); err != nil {
		return err
	}

	isCompleted := Todos[index].Completed

	if isCompleted {
		completionTime := time.Now()
		Todos[index].CompletedAt = &completionTime
	}

	Todos[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	Todos := *todos
	if err := Todos.validateIndex(index); err != nil {
		return err
	}

	Todos[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("ID", "Title", "Completed", "Created At", "Completed At")

	for index, todo := range *todos {
		Completed := "❌"
		CompletedAt := ""
		if todo.Completed {
			Completed = "✅"
			if todo.CompletedAt != nil {
				CompletedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), todo.Title, Completed, todo.CreatedAt.Format(time.RFC1123), CompletedAt)
	}

	table.Render()
}