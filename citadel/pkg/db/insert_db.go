package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InsertBook(name, author string) {
	db, err := sql.Open("mysql", Dsn())
	if err != nil {
		log.Printf("Unable to open Db connection to add book details \n")
		panic(err)
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// Query
	query := fmt.Sprintf("INSERT INTO book(book_name, book_author) VALUES('%s','%s');", name, author)

	_, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Unable to execute query to insert Book detail in db \n")
		panic(err)
	}
	log.Printf("Book added successfully \n")
}
