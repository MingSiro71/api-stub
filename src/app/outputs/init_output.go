package outputs

type InitOutput interface {
	Success()
	Error(string)
}
