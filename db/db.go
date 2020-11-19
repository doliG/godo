package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const dbPath = "/tmp/godo_db.json"

type Todo struct {
	Name    string
	Created time.Time /* Should be a date */
	Done    bool      /* Should be a date */
}

func GetAll() []Todo {
	if !dbExist() {
		createDb()
	}

	dat, err := ioutil.ReadFile(dbPath)
	if err != nil {
		fmt.Println("error", err.Error())
		panic(err)
	}
	var items []Todo
	err = json.Unmarshal(dat, &items)
	if err != nil {
		panic(err)
	}

	return items
}

func dbExist() bool {
	_, err := ioutil.ReadFile(dbPath) // Should use open instead and use proper error returned (ie. erro.code)
	if err != nil {
		return false
	}
	return true
}

func createDb() {
	fmt.Println("No db file found. Creating it at ", dbPath)
	var items []Todo
	jsoned, err := json.Marshal(items)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(dbPath, jsoned, 0644)
	if err != nil {
		panic(err)
	}
}

func Add(item Todo) {
	items := GetAll()
	items = append(items, item)

	UpdateAll(items)
}

func UpdateAll(items []Todo) {
	jsoned, _ := json.Marshal(items)
	err := ioutil.WriteFile(dbPath, jsoned, 0644)
	if err != nil {
		panic(err)
	}
}
