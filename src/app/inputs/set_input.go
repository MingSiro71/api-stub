package inputs

import (
	"api_stub/dtos"
)

type SetInput interface {
	Handle(dtos.SetDto)
}
