package system

import (
	"github.com/go-glx/input/system/action"
	"github.com/go-glx/input/system/controller"
	"github.com/go-glx/input/system/data"
	"github.com/go-glx/input/system/internal/binding"
)

func BindAs[T data.Type](
	sys *System,
	ctl *controller.Controller,
	action *action.Action,
	binders ...binding.F[T],
) {
	for _, binder := range binders {
		bind(sys, action.ID(), ctl, binder)
	}
}

func bind[T data.Type](s *System, actionID action.ID, ctl *controller.Controller, binder binding.F[T]) {
	if s.binders == nil {
		s.binders = make(map[action.ID]map[controller.ID][]binding.F[any])
	}

	if _, exist := s.binders[actionID]; !exist {
		s.binders[actionID] = make(map[controller.ID][]binding.F[any])
	}

	if _, exist := s.binders[actionID][ctl.ID()]; !exist {
		s.binders[actionID][ctl.ID()] = make([]binding.F[any], 0)
	}

	s.binders[actionID][ctl.ID()] = append(s.binders[actionID][ctl.ID()], func() any {
		return binder()
	})
}
