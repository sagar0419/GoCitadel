package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CheckDb(bookname string) string {
	query := fmt.Sprintf("SELECT book_author FROM book WHERE book_name REGEXP '%s';", bookname)
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	db, err := sql.Open("mysql", Dsn())
	if err != nil {
		log.Printf("Unable to open DB to check book record \n")
		panic(err)
	}
	defer db.Close()
	var Bookrecord string
	err = db.QueryRowContext(ctx, query).Scan(&Bookrecord)
	if err != nil {
		log.Printf("Unable to exec query on db to check book record \n")
		return ""
	}
	return Bookrecord
}
