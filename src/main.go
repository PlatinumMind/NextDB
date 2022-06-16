package main

import (
	// "fmt"
	"os"
	"log"

	// "nextdb.dev/src/database"
	"nextdb.dev/src/parser"
)

type Data struct {
	name string
	age int
	grade string
}

func main() {
	/*db := database.NewNextDatabase("mydb", "root", "password", 9903)
	db.NewCell("test")
 	db.NewDocument(Data{"arif", 12, "2 year middleschool"}, "test")
 	fmt.Printf("db: %v\n", db)*/

	_, err := os.ReadFile("./test.dql");
	if err != nil {
		log.Fatal(err)
	}
	
	parser.ParseDQL("create database \"test1\" set username \"arif\" set password \"pass\"")
}