package parser

import (
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

type Stated struct {
	Machine string //Unique ID for the machine (Eg:- FQDN)
	Start   []string
	File    []file
	// 	Package package
	// 	Service service
	// 	Exec exec
	// 	User user
	// 	Group group
	// 	Schedule schedule
	External []external
}

type file struct {
	Name        string
	Destination string
	Source      string
	Mode        os.FileMode
}

type external struct {
	Name   string
	Plugin string
	Config string
}

func Parse(state_file_path string) (Stated, error) {
	var stated Stated
	if _, err := toml.DecodeFile(state_file_path, &stated); err != nil {
		log.Println(err)
		return stated, err
	}
	return stated, nil
}
