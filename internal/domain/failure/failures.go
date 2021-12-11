package failure

const (
	noPhone = "phone num to dede bro"
	noEmail = "yrr email kon dega"
)

func NoPhone() error {
	return failure{msg: noPhone}
}

func NoEmail() error {
	return failure{msg: noEmail}
}
