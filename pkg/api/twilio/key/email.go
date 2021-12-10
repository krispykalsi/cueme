package key

import "os"

func EmailAuth() string {
	return os.Getenv("TWILIO_EMAIL_AUTH_KEY")
}

func FromVerifiedEmail() string {
	return os.Getenv("TWILIO_VERIFIED_EMAIL")
}
