package handlers

import (
	"cueme/internal/domain/enum/medium"
	"cueme/internal/domain/failure"
	"cueme/internal/domain/model"
	"cueme/pkg/api/twilio"
	"github.com/gin-gonic/gin"
	"time"
)

func handleEndpoint(c *gin.Context, req *model.MsgRequest, m medium.Medium) {
	cueTime, parseErr := req.GetCueTime()
	if parseErr != nil {
		handleError(c, parseErr)
		return
	}

	if ok, dataErr := isAllDataValid(req, m); ok {
		now := time.Now()
		if cueTime.After(now) {
			time.AfterFunc(cueTime.Sub(now), func() {
				_ = callApi(req, m)
			})
		} else {
			apiErr := callApi(req, m)
			handleError(c, apiErr)
		}
	} else {
		handleError(c, dataErr)
	}
}

func callApi(req *model.MsgRequest, m medium.Medium) error {
	var err error

	switch m {
	case medium.Sms:
		err = twilio.SendViaSms(req.Phone, req.Message)
	case medium.Whatsapp:
		err = twilio.SendViaWhatsapp(req.Phone, req.Message)
	case medium.Email:
		err = twilio.SendViaEmail(req.Email, req.Message)
	}

	return err
}

func isAllDataValid(req *model.MsgRequest, m medium.Medium) (bool, error) {
	isPhoneOk := req.HasValidPhone()
	isEmailOk := req.HasValidEmail()
	var err error

	if (m == medium.Sms || m == medium.Whatsapp) && !isPhoneOk {
		err = failure.NoPhone()
	} else if m == medium.Email && !isEmailOk {
		err = failure.NoEmail()
	}

	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}
