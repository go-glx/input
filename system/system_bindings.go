package system

import (
	"github.com/go-glx/input/system/controller"
	"github.com/go-glx/input/system/data"
	"github.com/go-glx/input/system/internal/binding"
)

func DefineActionOf[T data.Type](s *System, action actionID, ctl *controller.Controller, binder ...binding.F[T]) {
	for _, binder := range binder {
		bindAction(s, action, ctl, binder)
	}
}

func bindAction[T data.Type](s *System, action actionID, ctl *controller.Controller, binder binding.F[T]) {
	if s.binders == nil {
		s.binders = make(map[actionID]map[controller.ID][]binding.F[any])
	}

	if _, exist := s.binders[action]; !exist {
		s.binders[action] = make(map[controller.ID][]binding.F[any])
	}

	if _, exist := s.binders[action][ctl.ID()]; !exist {
		s.binders[action][ctl.ID()] = make([]binding.F[any], 0)
	}

	s.binders[action][ctl.ID()] = append(s.binders[action][ctl.ID()], func() any {
		return binder()
	})
}
