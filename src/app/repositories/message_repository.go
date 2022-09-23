package repositories

import (
	"api_stub/vo"
)

type MessageRepository interface {
	Push(vo.Id, string) error
	Pop(vo.Id) (string, error)
	// list(vo.Id) ([]string, error)
	// clear(vo.Id) error
	// init() error
}
