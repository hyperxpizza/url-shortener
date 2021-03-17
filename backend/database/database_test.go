package database

import "testing"

var database Database

func init() {
	database = NewDatabase()
}

func TestExistsID(t *testing.T) {
	database.existsID(1)
}
