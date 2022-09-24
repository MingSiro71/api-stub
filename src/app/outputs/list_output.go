package outputs

import (
	"api_stub/vo"
)

type ListOutput interface {
	Success(vo.Id, []string)
}
