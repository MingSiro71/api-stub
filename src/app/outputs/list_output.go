package outputs

type ListOutput interface {
	Success([]string)
	Error(string)
}
