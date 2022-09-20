package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func main {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.Default())
	router.GET("/api/entries", getEntries)
	router.POST("/api/entries", postEntry)
	router.GET("/api/entries/:id", getEntry)
	router.PUT("/api/entries/:id", putEntry)
	router.DELETE("/api/entries/:id", deleteEntry)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
		
}