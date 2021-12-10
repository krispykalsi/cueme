package main

import (
	"cueme/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func initialiseRoutes(g *gin.Engine) {
	g.POST("/api/sms", handlers.Sms)
	g.POST("/api/wa", handlers.Whatsapp)
	g.POST("/api/email", handlers.Email)
}

func runServer() {
	r := gin.Default()
	initialiseRoutes(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
