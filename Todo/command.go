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
	CmdFlags := cmdFlags{}

	flag.StringVar(&CmdFlags.add, "add", "", "Add a new todo specify title")
	flag.StringVar(&CmdFlags.edit, "edit", "", "Edit a todo by index and specify new title. id:new_title")
	flag.IntVar(&CmdFlags.delete, "delete", -1, "Specify todo by index to delete")
	flag.IntVar(&CmdFlags.toggle, "toggle", -1, "Specify todo by index to toggle")
	flag.BoolVar(&CmdFlags.list, "list", false, "List all todos")

	flag.Parse()

	return &CmdFlags
}

func (CmdFlags *cmdFlags) Execute(todos *Todos) {
	switch {
		case CmdFlags.list:
			todos.print()
		
		case CmdFlags.add != "":
			todos.Add(CmdFlags.add)
		
		case CmdFlags.edit != "":
			parts := strings.SplitN(CmdFlags.edit, ":", 2)
			if len(parts) != 2 {
				fmt.Println("Error, invalid format for edit. Please use id:new_title")

				index, err := strconv.Atoi(parts[0])

				if err != nil {
					fmt.Println("Error, invalid index for edit. Please use id:new_title")
					os.Exit(1)
				}

				todos.Edit(index, parts[1])
			}
		
		case CmdFlags.toggle != -1:
			todos.Toggle(CmdFlags.toggle)

		case CmdFlags.delete != -1:
			todos.delete(CmdFlags.delete)

		default:
			fmt.Println("Invalid Command")
	}
}