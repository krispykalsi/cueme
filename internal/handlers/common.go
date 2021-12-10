package handlers

import (
	"cueme/internal/domain/enum/medium"
	"cueme/internal/domain/failure"
	"cueme/internal/domain/model"
	"cueme/pkg/api/twilio"
	"github.com/gin-gonic/gin"
)

func handleError(c *gin.Context, err error) {
	if err != nil {
		sendBadResponse(c, err.Error())
	} else {
		sendOKResponse(c)
	}
}

func handleEndpoint(c *gin.Context, req *model.MsgRequest, m medium.Medium) {
	isPhoneOk := req.HasValidPhone()
	isEmailOk := req.HasValidEmail()
	var err error

	if m == medium.Sms {
		if isPhoneOk {
			err = twilio.SendViaSms(req.Phone, req.Message)
		} else {
			err = failure.NoPhone()
		}
	} else if m == medium.Whatsapp {
		if isPhoneOk {
			err = twilio.SendViaWhatsapp(req.Phone, req.Message)
		} else {
			err = failure.NoPhone()
		}
	} else if m == medium.Email {
		if isEmailOk {
			err = twilio.SendViaEmail(req.Email, req.Message)
		} else {
			err = failure.NoEmail()
		}
	} else {
		err = failure.BadMedium()
	}

	handleError(c, err)
}
