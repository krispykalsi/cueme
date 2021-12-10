package handlers

import (
	"cueme/internal/domain/model"
	"github.com/gin-gonic/gin"
)

func parseRequestToJSON(c *gin.Context) *model.MsgRequest {
	req := &model.MsgRequest{}
	err := c.BindJSON(req)
	if err != nil {
		sendBadResponse(c, "bad format")
	}
	return req
}

func sendBadResponse(c *gin.Context, msg string) {
	c.JSON(502, gin.H{"err": msg})
}

func sendOKResponse(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "ok, got it",
	})
}
