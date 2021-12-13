package handlers

import (
	"cueme/internal/domain/enum/medium"
	"github.com/gin-gonic/gin"
)

func Sms(c *gin.Context) {
	req := parseRequestToJSON(c)
	handleEndpoint(c, req, medium.Sms)
}

func Whatsapp(c *gin.Context) {
	req := parseRequestToJSON(c)
	handleEndpoint(c, req, medium.Whatsapp)
}

func Email(c *gin.Context) {
	req := parseRequestToJSON(c)
	handleEndpoint(c, req, medium.Email)
}

func Api(c *gin.Context) {
	type json map[string]interface{}

	c.IndentedJSON(200, json{
		"info": "Welcome to the CuemeApi",
		"routes": json{
			"whatsapp": "cueme.xyz/api/wa",
			"sms":      "cueme.xyz/api/sms",
			"email":    "cueme.xyz/api/email",
		},
		"schema": json{
			"message": "hello cueme :3",
			"phone":   "+919876543210",
			"email":   "joe@mama.com",
			"time":    "2021-07-25T15:59:59.000Z",
		},
	})
}
