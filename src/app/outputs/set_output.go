package outputs

import (
	"api_stub/vo"
)

type SetOutput interface {
	Success(vo.Id)
	Error(string)
}
