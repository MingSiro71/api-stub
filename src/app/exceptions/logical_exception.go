package exceptions

type LogicalException struct {
	msg string
}

const LogicalExceptionDefault = "Internal server Exception."

func NewLogicalException(m string) *LogicalException {
	return &LogicalException{msg: m}
}

func (err *LogicalException) Error() string {
	return err.msg
}
