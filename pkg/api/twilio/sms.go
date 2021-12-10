package twilio

import (
	"cueme/pkg/api/twilio/key"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const smsEndpoint = key.MessagesApiEndpoint

func SendViaSms(phone string, msg string) error {
	data := url.Values{}
	data.Set("To", phone)
	data.Set("MessagingServiceSid", key.MessagingServiceSid)
	data.Set("Body", msg)

	client := &http.Client{}
	r, err := http.NewRequest("POST", smsEndpoint, strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.SetBasicAuth(key.PhoneAuthUsername, key.PhoneAuthPassword)

	return executeRequest(client, r)
}
