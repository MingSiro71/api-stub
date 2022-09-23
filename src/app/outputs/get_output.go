package outputs

type GetOutput interface {
	Success(string)
	Error(string)
}
