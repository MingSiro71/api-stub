package exceptions

type RoutingException struct {
	msg string
}

const RoutingExceptionDefault = "Not found."

func NewRoutingException(m string) *RoutingException {
	return &RoutingException{msg: m}
}

func (err *RoutingException) Error() string {
	return err.msg
}
