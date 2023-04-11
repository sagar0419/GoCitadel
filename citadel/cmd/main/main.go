package main

import (
	books "citadel/api/books"
	homepage "citadel/api/homepage"
	record "citadel/api/record"
	database "citadel/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
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

	// Server
	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
