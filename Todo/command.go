package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	add    string
	delete int
	edit   string
	toggle int
	list   bool
}

func NewCmdFlags() *cmdFlags {
	CF := cmdFlags{}

	flag.StringVar(&CF.add, "add", "", "Add a new todo specify title")
	flag.StringVar(&CF.edit, "edit", "", "Edit a todo by index and specify new title. id:new_title")
	flag.IntVar(&CF.delete, "delete", -1, "Specify todo by index to delete")
	flag.IntVar(&CF.toggle, "toggle", -1, "Specify todo by index to toggle")
	flag.BoolVar(&CF.list, "list", false, "List all todos")

	flag.Parse()

	return &CF
}

func (CF *cmdFlags) Execute(todos *Todos) {
	switch {
		case CF.list:
			todos.print()
		
		case CF.add != "":
			todos.Add(CF.add)
		
		case CF.edit != "":
			parts := strings.SplitN(CF.edit, ":", 2)
			if len(parts) != 2 {
				fmt.Println("Error, invalid format for edit. Please use id:new_title")

				index, err := strconv.Atoi(parts[0])

				if err != nil {
					fmt.Println("Error, invalid index for edit. Please use id:new_title")
					os.Exit(1)
				}

				todos.Edit(index, parts[1])
			}
		
		case CF.toggle != -1:
			todos.Toggle(CF.toggle)

		case CF.delete != -1:
			todos.delete(CF.delete)

		default:
			fmt.Println("Invalid Command")
	}
}