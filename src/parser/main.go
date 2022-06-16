package parser

import (
	"fmt"
	"strings"
	// "nextdb.dev/src/database"
)

func ParseDQL(query string) {
	lines := strings.Split(query, ";")
	for _, line := range lines {
		data := strings.Split(line, " ")
		if data[0] == "create" {
			if data[1] == "database" {
				if data[3] == "set" {
					if data[4] == "password" {
						var tmp = data[8]
						fmt.Println(tmp)
						fmt.Println("db name: "+ data[2] + " user: " + data[8] + " passowrd: " + data[5] + " port: " + string(rune(9954)))
						// database.NewNextDatabase(data[2], data[8], data[5], 9954)
					}		

					if data[4] == "username" {
						fmt.Println("db name: "+data[2]+"\nuser: "+data[5]+"\npassowrd: "+data[8]+"\nport: ", 9954)
						// database.NewNextDatabase(data[2], data[5], data[8], 9954)
					}
				}
			}

			if data[1] == "" {
				
			}
		}
	}
}