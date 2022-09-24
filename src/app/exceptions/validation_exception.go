package exceptions

type ValidationException struct {
	msg string
}

const ValidationExceptionDefault = "Invalid request."

func NewValidationException(m string) *ValidationException {
	return &ValidationException{msg: m}
}

func (err *ValidationException) Error() string {
	return err.msg
}
