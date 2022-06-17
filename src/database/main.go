package database

import (
	"encoding/json"
	"errors"
	"io"

	// "fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"github.com/lucsky/cuid"
)


	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func NewNextDatabase(name string, username string, password string, port uint32) {
	if err := os.Mkdir("/etc/nextdb/", 0755); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Panic(err)
		}
	}
	if err := os.Mkdir("/etc/nextdb/databases", 0755); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Panic(err)
		}
	}
	if err := os.Mkdir("/etc/nextdb/databases/" + name, 0755); err != nil {
		if !errors.Is(err, os.ErrExist) {
			log.Panic(err)
		}
	}

	conf := &NextDatabase{
		name, 
		username, 
		password, 
		int32(port),
	}

	data, _ := yaml.Marshal(conf)
	err := os.WriteFile("/etc/nextdb/databases/" + name +"/database.yaml", data, 0644)
	check(err) 
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