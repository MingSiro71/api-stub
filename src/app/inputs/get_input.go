package inputs

import (
	"api_stub/dtos"
)

type GetInput interface {
	Handle(dtos.QueryDto) error
}
