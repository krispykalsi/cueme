package twilio

import (
	"cueme/pkg/api/twilio/key"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const emailEndpoint = "https://api.sendgrid.com/v3/mail/send"

func SendViaEmail(email string, msg string) error {
	data := `{"personalizations": [{"to": [{"email": "` + email + `"}]}],"from": {"email": "` + key.FromVerifiedEmail + `"},"subject": "Message from Cueme","content": [{"type": "text/plain", "value": "` + msg + `"}]}`

	client := &http.Client{}
	r, err := http.NewRequest("POST", emailEndpoint, strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Authorization", "Bearer "+key.EmailAuthKey)
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Content-Length", strconv.Itoa(len(data)))

	return executeRequest(client, r)
}
