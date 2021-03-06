package main

import (
	"cueme/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func initialiseRoutes(g *gin.Engine) {
	g.GET("/api", handlers.Api)
	g.POST("/api/sms", handlers.Sms)
	g.POST("/api/wa", handlers.Whatsapp)
	g.POST("/api/email", handlers.Email)
}

func main() {
	r := gin.Default()
	initialiseRoutes(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
