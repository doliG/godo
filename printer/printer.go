package printer

import (
	"fmt"

	"github.com/doliG/godo/db"
	"github.com/gookit/color"
)

func PrintAll(items []db.Todo, printAll bool) {
	for index, item := range items {
		if item.Done && !printAll {
			continue
		}
		Print(index, item)
	}
}

func Print(index int, item db.Todo) {
	if item.Done {
		color.C256(244).Printf("%d. [✔] %s\n", index, item.Name)
	} else {
		fmt.Printf("%d. [ ] %s\n", index, item.Name)
	}
}
