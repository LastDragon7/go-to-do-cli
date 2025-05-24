package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new values to specify the title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit the title of the given index")
	flag.IntVar(&cf.Del, "del", -1, "Delete the task of the given index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle the task of the given index")
	flag.BoolVar(&cf.List, "list", false, "Display all the Tasks")

	flag.Parse()

	return &cf

}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitAfterN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error , invalid input Format. Use id:new_title")
			os.Exit(1)
		}
		idx, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error : invalid index for edit")
			os.Exit(1)
		}

		todos.edit(idx, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid Command")
	}
}
