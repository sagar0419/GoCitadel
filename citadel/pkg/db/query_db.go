package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func QueryDb() ([]Book, error) {
	db, err := sql.Open("mysql", Dsn())
	if err != nil {
		log.Printf("Unable to open db connection to quer db \n")
		panic(err)
	}
	defer db.Close()

	//DB Query
	query := `SELECT book_name, book_author FROM book;`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		log.Print("unable to prepare statement for query on books db \n ")
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Printf("unable to execute query on book db \n")
	}

	defer rows.Close()

	var details []Book
	for rows.Next() {
		var b Book
		err := rows.Scan(&b.Title, &b.Author)
		if err != nil {
			log.Print("unable to scan row in db", err)
			return nil, err
		}
		details = append(details, b)
	}
	return details, nil
}
