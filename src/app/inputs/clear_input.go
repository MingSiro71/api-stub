package inputs

import (
	"api_stub/dtos"
)

type ClearInput interface {
	Handle(dtos.QueryDto) error
}
