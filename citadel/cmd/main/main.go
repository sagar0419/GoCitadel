package main

import (
	books "citadel/api/books"
	healthz "citadel/api/healthz"
	homepage "citadel/api/homepage"
	record "citadel/api/record"
	database "citadel/pkg/db"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	go func() {
		// Simulate app boot: load config, connect to DB, etc.
		time.Sleep(5 * time.Second)

		// After all startup is complete
		healthz.MarkInitialized()
		log.Println("✅ App initialized")
	}()

	database.CreateDb()
	database.CreateTable()
	router := gin.Default()

	// Trusted Proxy
	router.SetTrustedProxies([]string{"127.0.0.1"})

	// HomePage
	router.GET("/", homepage.HomePage)

	// Storing Book Info
	router.GET("/books", books.Books)

	// Storing Book Info
	router.GET("/record", record.BookDetail)

	// Health Check
	router.GET("/healthz", healthz.Healthz)  // Browsers use GET by default.
	router.HEAD("/healthz", healthz.Healthz) // Curl uses HEAD if you do something like: curl -I http://localhost:3000/healthz

	// Server
	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
