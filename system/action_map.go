package system

import (
	"reflect"

	"github.com/go-glx/input/system/binding"
	"github.com/go-glx/input/system/data"
)

type (
	ActionMapID string

	ActionMap struct {
		ID      ActionMapID
		Actions []*Action[data.Type]
	}
)

func NewActionMap(name ActionMapID) *ActionMap {
	return &ActionMap{
		ID:      name,
		Actions: make([]*Action[data.Type], 0),
	}
}

func BindTo[T data.Type](aMap *ActionMap, actionID ActionID, bindings []*binding.Binding) {
	action := NewAction[T](actionID, bindings...)
	aMap.Actions = append(aMap.Actions, reflect.ValueOf(action).Interface().(*Action[data.Type]))
}
