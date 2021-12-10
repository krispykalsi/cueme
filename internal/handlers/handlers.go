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
