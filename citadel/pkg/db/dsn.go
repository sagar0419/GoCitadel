package db

import "fmt"

const (
	username = "root"
	password = "password"
	hostname = "127.0.0.1:3306"
	dbname   = "citadel"
	// olddbname = "db"
)

// Creating mysql connection string "user:password@tcp(127.0.0.1:3306)/database_name"

func CreatDbDsn() string {
	// return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, olddbname)
	return fmt.Sprintf("%s:%s@tcp(%s)/", username, password, hostname)
}

func Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}
