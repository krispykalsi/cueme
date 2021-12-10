package failure

const (
	noPhone    = "phone num to dede bro"
	noEmail    = "yrr email kon dega"
	badMedium  = "wrong medium. Only choices are sms / phone / whatsapp"
)

func NoPhone() error {
	return failure{msg: noPhone}
}

func NoEmail() error {
	return failure{msg: noEmail}
}

func BadMedium() error {
	return failure{msg: badMedium}
}
