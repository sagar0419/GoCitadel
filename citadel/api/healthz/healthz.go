package healthz

import (
	data "citadel/pkg/db"
	"database/sql"
	"log"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
)

var initialized atomic.Bool

// Function to call when the app finishes initialization
func MarkInitialized() {
	initialized.Store(true)
}

// Flag to check inside health check
func IsInitialized() bool {
	return initialized.Load()
}

func Healthz(c *gin.Context) {

	if !IsInitialized() {
		c.String(http.StatusServiceUnavailable, "Service is starting up")
		return
	}

	// Checking MySql connectivity
	db, err := sql.Open("mysql", data.Dsn())
	if err != nil {
		log.Printf("Unable to open db connection to query db \n")
		c.String(http.StatusInternalServerError, "Database connection failed")
		panic(err)
	}
	defer db.Close()

	// Ping the database to check connectivity
	if err := db.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
		c.String(http.StatusServiceUnavailable, "Database unreachable")
		panic(err)
	}

	c.Header("Content-Type", "text/plain")
	message := ("Service is up and runing")
	c.String(200, message)
}
