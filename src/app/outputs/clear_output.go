package outputs

import (
	"api_stub/vo"
)

type ClearOutput interface {
	Success(vo.Id, int)
}
