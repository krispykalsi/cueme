package twilio

import (
	"cueme/pkg/api/twilio/key"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func SendViaWhatsapp(phone string, msg string) error {
	data := url.Values{}
	data.Set("To", "whatsapp:"+phone)
	data.Set("From", "whatsapp:"+key.WhatsappPhoneNumber())
	data.Set("Body", msg)

	client := &http.Client{}
	r, err := http.NewRequest("POST", key.MessagesApiEndpoint(), strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	r.SetBasicAuth(key.PhoneAuthUsername(), key.PhoneAuthPassword())

	return executeRequest(client, r)
}
