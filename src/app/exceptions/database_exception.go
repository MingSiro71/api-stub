package exceptions

type DatabaseException struct {
	msg string
}

const DatabaseExceptionDefault = "Internal server Exception."

func NewDatabaseException(m string) *DatabaseException {
	return &DatabaseException{msg: m}
}

func (err *DatabaseException) Error() string {
	return err.msg
}
