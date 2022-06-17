package parser

import (
	// "fmt"
	"strings"
	"nextdb.dev/src/database"
)

func trimQuotes(s string) string {
    if len(s) >= 2 {
        if s[0] == '"' && s[len(s)-1] == '"' {
            return s[1 : len(s)-1]
        }
    }
    return s
}



func ParseDQL(query string) {
	lines := strings.Split(query, ";")
	for _, line := range lines {
		data := strings.Split(line, " ")
		if data[0] == "create" {
			if data[1] == "database" {
				if data[3] == "set" {
					if data[4] == "password" {
						// fmt.Println("db name: "+ trimQuotes(data[2]) + " user: " + trimQuotes(data[8]) + " passowrd: " + trimQuotes(data[5]) + " port: " + string(rune(9954)))
						database.NewNextDatabase(data[2], data[8], data[5], 9954)
					}		

					if data[4] == "username" {
						// fmt.Println("db name: "+trimQuotes(data[2])+"\nuser: "+trimQuotes(data[5])+"\npassowrd: "+trimQuotes(data[8])+"\nport: ", 9954)
						database.NewNextDatabase(trimQuotes(data[2]), trimQuotes(data[5]), trimQuotes(data[8]), 9954)
					}
				}
			}

			if data[1] == "cell" {
				
			}
		}
	}
}