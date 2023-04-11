package record

import (
	"citadel/pkg/db"
	"fmt"

	"github.com/gin-gonic/gin"
)

func BookDetail(c *gin.Context) {
	list, _ := db.QueryDb()
	for _, details := range list {
		BookInfo := fmt.Sprintf("Book name is %v and its author is %v \n", details.Title, details.Author)
		c.Header("Content-Type", "text/plain")
		c.String(200, BookInfo)
	}
}
