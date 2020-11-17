package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gookit/color"
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
		fmt.Println("subcommand 'add'")
		fmt.Println("  tail:", addCmd.Args())

	case "list":
		listCmd.Parse(os.Args[2:])
		list(*listAll)

	default:
		fmt.Println("expected subcommands. Type ", os.Args[0], "help", "for help")
		os.Exit(1)
	}
}

type todo struct {
	name    string
	created string /* Should be a date */
	done    bool   /* Should be a date */
}

var items = []todo{
	{name: "Make a cake", created: "17 nov 11:15", done: false},
	{name: "Do pushups", created: "15 nov 10:32", done: true},
	{name: "Buy Nintendo switch controller", created: "15 nov 10:32", done: false},
}

/**
 * List
 */
func list(listAll bool) {
	for _, item := range items {
		if item.done {
			if listAll {
				color.Gray.Println("- [x]", item.created, ":", item.name)
			}
		} else {
			fmt.Println("- [ ]", item.created, ":", item.name)
		}
	}
}
