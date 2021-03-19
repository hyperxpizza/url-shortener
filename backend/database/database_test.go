package database

import (
	"fmt"
	"log"
	"testing"
)

var database Database

func init() {
	database = NewDatabase()
}

func TestExistsID(t *testing.T) {
	e := database.IDexists(666)
	fmt.Println(e)
}

func TestInsertItem(t *testing.T) {
	id, err := database.Insert("https://google.com", 0)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("encoded id: ", id)
}

func TestLoad(t *testing.T) {
	s, err := database.Get("O8KEZlAseeb")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)
}
