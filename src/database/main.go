package database

import (
	"encoding/json"
	"errors"
	"io"

	// "fmt"
	"log"
	"os"

	"github.com/lucsky/cuid"
)


func NewNextDatabase(name string, username string, password string, port uint32) *NextDatabase {
	if err := os.Mkdir("/etc/nextdb/databases/" + name, 0755); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Panic(err)
		}
	}
	return &NextDatabase{name, username, password, int32(port)}
}


func (d *NextDatabase) NewDocument(data interface{}, cell string) {
	file, err := json.Marshal(data)
	if err != nil {
        log.Println(err)
        return
    }
	// log.Println("/etc/nextdb/database/" + d.name + "/cells/" + cell + "/" + cuid.New() + ".json")
	
	if _file, err := os.Create("/etc/nextdb/database/" + d.name + "/cells/" + cell + "/" + cuid.New() + ".json"); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Panic(err)
		}
		io.WriteString(_file, string(file))
		defer _file.Close()
		// _file.WriteString(string(file))
	}
	// file = append(file, )
	// if err := os.WriteFile("/etc/nextdb/database/" + d.name + "/cells/" + cell + "/" + cuid.New() + ".json", []byte(file), 0755); err != nil {
	// 	if !errors.Is(err, os.ErrExist) {
	// 		log.Panic(err)
	// 	}
	// }
}

func (d *NextDatabase) NewCell(cellName string) {
	os.Mkdir("/etc/nextdb/databases/" + d.name + "/cells/", 0755)
	if err := os.Mkdir("/etc/nextdb/databases/" + d.name + "/cells/" + cellName, 0755); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Panic(err)
		}
	}	
}