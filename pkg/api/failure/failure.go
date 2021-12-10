package failure

const unexpected = "something went wrong? (panic noises)"

type failure struct {
	msg string
}

func (e failure) Error() string {
	if e.msg != "" {
		return e.msg
	} else {
		return unexpected
	}
}
