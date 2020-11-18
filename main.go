package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/doliG/godo/db"
	"github.com/doliG/godo/printer"
)

func main() {
	// https://gobyexample.com/command-line-subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listAll := listCmd.Bool("all", false, "Display all tasks, including done ones.")

	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		name := strings.Join(addCmd.Args(), " ")
		add(name)

	case "list":
		listCmd.Parse(os.Args[2:])
		list(*listAll)

	default:
		fmt.Println("expected subcommands. Type ", os.Args[0], "help", "for help")
		os.Exit(1)
	}
}

func list(listAll bool) {
	todos := db.GetAll()

	visibleItems := []db.Todo{}
	for _, item := range todos {
		if item.Done && !listAll {
			continue
		}
		visibleItems = append(visibleItems, item)
	}

	printer.PrintAll(visibleItems)
}

func add(name string) {
	item := db.Todo{
		Name:    name,
		Created: "Now",
		Done:    false,
	}
	db.Add(item)
}
