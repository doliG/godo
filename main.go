package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/doliG/welldone/db"
	"github.com/doliG/welldone/printer"
	"github.com/gookit/color"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		addCmd.Parse(os.Args[2:])
		name := strings.Join(addCmd.Args(), " ")
		add(name)

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		listAll := listCmd.BoolP("all", "a", false, "List all tasks, including done ones.")
		listCmd.Parse(os.Args[2:])
		query := listCmd.Arg(0)
		list(query, *listAll)

	case "toggle":
		toggleCmd := flag.NewFlagSet("toggle", flag.ExitOnError)
		toggleCmd.Parse(os.Args[2:])
		ids := toggleCmd.Args()
		toggle(ids)

	case "edit":
		editCmd := flag.NewFlagSet("edit", flag.ExitOnError)
		editMessage := editCmd.StringP("message", "m", "", "New name for your task")
		editCmd.Parse(os.Args[2:])
		id := editCmd.Arg(0)
		edit(id, *editMessage)

	case "version":
		version()

	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Printf(`Welldone is a CLI todolist.
Usage:
	list     [query] [--all, -a]         List todos
	add      [name]                      Add a todo
	edit     [id] [--message, -m name]   Edit the name of the todo with the given id
	toggle   [id [id2, id3...]]          Mark an idea as done if its not, undone otherwise
	version                              Print the current version
`)
	os.Exit(0)
}

func list(query string, listAll bool) {
	todos := db.GetAll()
	toShow := todos

	if len(query) > 0 {
		toShow = []db.Todo{}
		for _, item := range todos {
			if strings.Contains(item.Name, query) {
				toShow = append(toShow, item)
			}
		}
	}

	printer.PrintAll(toShow, listAll)
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

func edit(id string, newName string) {
	todos := db.GetAll()

	// TODO: Maybe it can be done by CLI parser ?
	index, err := strconv.Atoi(id)
	if err != nil {
		color.Warn.Tips("Cannot convert '%s' into number. Skipping...", id)
	} else if index < 0 || index >= len(todos) {
		color.Warn.Tips("Invalid id '%s'. It must be  >0 and <%d. Skipping...", id, len(todos))
	}
	if len(newName) == 0 {
		color.Error.Println("You must specify a new name with -m option.")
		os.Exit(1)
	}

	todos[index].Name = newName
	db.UpdateAll(todos)
}

func version() {
	progname := "welldone"
	version := "0"
	fmt.Printf("%s v%s\n", progname, version)
}
