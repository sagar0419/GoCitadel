package test

import (
	"citadel/pkg/db"
	"log"
	"testing"
)

func TestBook(t *testing.T) {
	bookname := "Go Lang lib test"
	bookauthor := "tester"
	db.InsertBook(bookname, bookauthor)

	result := db.CheckDb(bookname)
	if result != "" {
		message := "Test passed"
		log.Printf(message)
		return
	} else {
		message := "Test Failed"
		log.Printf(message)
		t.Errorf("Error: %s", message)
	}
}
