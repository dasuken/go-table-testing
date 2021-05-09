package repository

import (
	"fmt"
	"log"
	"testing"
)

var repo Database

func TestMain(m *testing.M) {
	setup()

	m.Run()
}

func setup() {
	if integrateFlag {
		fmt.Printf("Integration True \n")
		var err error
		repo, err = NewDatabase("root", "mysql", "127.0.0.1", "3306", "todo", "?parseTime=true&loc=Japan")
		if err != nil {
			log.Fatal(err)
		}
	}
}