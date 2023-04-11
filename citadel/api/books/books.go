package controllers

import (
	"citadel/pkg/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Books(c *gin.Context) {
	url := c.Request.URL.Query()
	BookName := url["title"]
	BookAuthor := url["author"]
	result := db.CheckDb(BookName[0])
	if result != "" {
		message := "Books details are already exsist in the database"
		c.Header("Content-Type", "text/plain")
		c.String(200, message)
		return
	}
	db.InsertBook(BookName[0], BookAuthor[0])
	DetailSaved := fmt.Sprintf("Your books details are saved")
	c.Header("Content-Type", "text/plain")
	c.String(200, DetailSaved)
}
