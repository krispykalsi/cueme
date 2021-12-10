package failure

func FromTwilio(msg string) error {
	return failure{msg: msg}
}
