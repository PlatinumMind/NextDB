package main

import (
	// "fmt"
	"nextdb.dev/src/database"
)

type Data struct {
	name string
	age int
	grade string
}

func main() {
	db := database.NewNextDatabase("mydb", "root", "password", 9903)
	db.NewCell("test")
	db.NewDocument(Data{"arif", 12, "2 year middleschool"}, "test")
	// fmt.Printf("db: %v\n", db)
}