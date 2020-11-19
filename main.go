package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/doliG/godo/db"
	"github.com/doliG/godo/printer"
	"github.com/gookit/color"
)

func main() {
	// https://gobyexample.com/command-line-subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listAll := listCmd.Bool("a", false, "List all tasks, including done ones.")

	toggleCmd := flag.NewFlagSet("toggle", flag.ExitOnError)

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

	case "toggle":
		toggleCmd.Parse(os.Args[2:])
		ids := toggleCmd.Args()
		toggle(ids)

	default:
		fmt.Println("expected subcommands. Type ", os.Args[0], "help", "for help")
		os.Exit(1)
	}
}

func list(listAll bool) {
	todos := db.GetAll()
	printer.PrintAll(todos, listAll)
}

func add(name string) {
	item := db.Todo{
		Name:    name,
		Created: time.Now(),
		Done:    false,
	}
	db.Add(item)
}

func toggle(ids []string) {
	todos := db.GetAll()

	for _, id := range ids {
		index, err := strconv.Atoi(id)
		if err != nil {
			color.Warn.Tips("Cannot convert '%s' into number. Skipping...", id)
			continue
		} else if index < 0 || index >= len(todos) {
			color.Warn.Tips("Invalid id '%s'. It must be  >0 and <%d. Skipping...", id, len(todos))
			continue
		}
		todos[index].Done = !todos[index].Done
	}

	db.UpdateAll(todos)
}
