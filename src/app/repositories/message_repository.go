package repositories

import (
	"api_stub/vo"
)

type MessageRepository interface {
	Push(vo.Id, string) error
	Pop(vo.Id) (string, error)
	List(vo.Id) ([]string, error)
	// Clear(vo.Id) error
	// Init() error
}
