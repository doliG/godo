package db

// import "encoding/json"

// Todo item
type Todo struct {
	Name    string
	Created string /* Should be a date */
	Done    bool   /* Should be a date */
}

var items = []Todo{
	{"Make a cake", "17 nov 11:15", false},
	{"Do pushups", "15 nov 10:32", true},
	{"Buy Nintendo switch controller", "15 nov 10:32", false},
}

func GetAll() []Todo {
	// t1 := Todo{"Make a kladkaka", "17 nov 11:15", false}
	// fmt.Println(t1)

	// b, err := json.Marshal(items)
	// if err != nil {
	// 	color.Error.Println("Error")
	// }
	// fmt.Println("json.Valid(b) ->", json.Valid(b))
	// fmt.Println("string(b) ->", string(b))

	// var t2 []Todo
	// err = json.Unmarshal(b, &t2)
	// fmt.Println(t2)

	return items
}
