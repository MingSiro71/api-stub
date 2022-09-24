package exceptions

type DataException struct {
	msg string
}

const DataExceptionDefault = "Data Conflict."

func NewDataException(m string) *DataException {
	return &DataException{msg: m}
}

func (err *DataException) Error() string {
	return err.msg
}
