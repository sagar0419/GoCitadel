package main

import (
	books "citadel/api/books"
	healthz "citadel/api/healthz"
	homepage "citadel/api/homepage"
	prom "citadel/api/metrics"
	record "citadel/api/record"
	database "citadel/pkg/db"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go func() {
		// Simulate app boot: load config, connect to DB, etc.
		time.Sleep(5 * time.Second)

		// After all startup is complete
		healthz.MarkInitialized()
		log.Println("*** App initialized ***")
	}()

	// Database
	database.CreateDb()
	database.CreateTable()
	router := gin.Default()

	// Prometheus
	router.Use(prom.PrometheusMiddleware())

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

	// Metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Server
	err := router.Run(":3000")
	if err != nil {
		panic(err)
	}
}
