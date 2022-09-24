package exceptions

type AuthException struct {
	msg string
}

const AuthExceptionDefault = "Not allowed."

func NewAuthException(m string) *AuthException {
	return &AuthException{msg: m}
}

func (err *AuthException) Error() string {
	return err.msg
}
