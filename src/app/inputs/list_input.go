package inputs

import (
	"api_stub/dtos"
)

type ListInput interface {
	Handle(dtos.QueryDto) error
}
