package printer

import (
	"fmt"

	"github.com/doliG/godo/db"
	"github.com/gookit/color"
)

func PrintAll(items []db.Todo) {
	for _, item := range items {
		Print(item)
	}
}

func Print(item db.Todo) {
	created := item.Created.Format("02 Jan, 15:04")
	if item.Done {
		color.Gray.Println("- [x]", created, ":", item.Name)
	} else {
		fmt.Println("- [ ]", created, ":", item.Name)
	}
}
