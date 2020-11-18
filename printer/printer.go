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
	if item.Done {
		color.Gray.Println("- [x]", item.Created, ":", item.Name)
	} else {
		fmt.Println("- [ ]", item.Created, ":", item.Name)
	}
}
