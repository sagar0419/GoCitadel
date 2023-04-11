package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func CreateDb() {
	db, err := sql.Open("mysql", CreatDbDsn())
	if err != nil {
		log.Printf("Error while connecting Db for DB creation \n")
		panic(err)
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err = db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Printf("Error Creating DB \n")
		panic(err)
	}
	log.Printf("DB is created")
}

func CreateTable() {
	db, err := sql.Open("mysql", Dsn())
	if err != nil {
		log.Printf("Unable to connect DB to create Table \n")
		panic(err)
	}
	defer db.Close()
	//  DB Query
	query := `CREATE TABLE IF NOT EXISTS book(id int primary key auto_increment, book_name varchar(2048) , book_author varchar(200) );`

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		log.Printf("Unable to create table in DB \n")
		panic(err)
	}
	log.Printf("Table is created successfully \n")
}
