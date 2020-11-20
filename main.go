package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/doliG/godo/db"
	"github.com/doliG/godo/printer"
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

	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Godo Usage:")
	fmt.Println("\tlist   [query] [--all, -a]")
	fmt.Println("\tadd    [name]")
	fmt.Println("\tedit   [id] [--message, -m name]")
	fmt.Println("\ttoggle [id [id2, id3...]]")
}

func list(query string, listAll bool) {
	todos := db.GetAll()

	var toShow []db.Todo
	if len(query) > 0 {
		for _, item := range todos {
			if strings.Contains(item.Name, query) {
				toShow = append(toShow, item)
			}
		}
	} else {
		toShow = todos
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
