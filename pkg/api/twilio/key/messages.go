package key

import "os"

func PhoneAuthUsername() string {
	return os.Getenv("TWILIO_PHONE_AUTH_USERNAME")
}

func PhoneAuthPassword() string {
	return os.Getenv("TWILIO_PHONE_AUTH_PASSWORD")
}

func MessagingServiceSid() string {
	return os.Getenv("TWILIO_MSG_SERVICE_SID")
}

func WhatsappPhoneNumber() string {
	return os.Getenv("TWILIO_WA_PHONE_NO")
}

func MessagesApiEndpoint() string {
	return os.Getenv("TWILIO_MSG_API_ENDPOINT")
}
