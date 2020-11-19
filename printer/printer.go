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
	created := item.Created.Format("02 Jan, 15:04")
	if item.Done {
		color.Gray.Printf("%d. %s: (done) %s\n", index, created, item.Name)
	} else {
		fmt.Printf("%d. %s : %s\n", index, created, item.Name)
	}
}
