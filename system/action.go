package system

import (
	"github.com/go-glx/input/system/binding"
	"github.com/go-glx/input/system/data"
)

type (
	ActionKind uint8
	ActionID   any

	Action[T data.Type] struct {
		id       ActionID
		bindings []*binding.Binding
	}
)

func NewAction[T data.Type](id ActionID, bindings ...*binding.Binding) *Action[T] {
	return &Action[T]{
		id:       id,
		bindings: bindings,
	}
}
