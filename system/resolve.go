package system

import (
	"fmt"

	"github.com/go-glx/input/system/action"
	"github.com/go-glx/input/system/data"
)

// Resolve will fetch actual input values for provided action
// if you have more than 1 player, you need switch it with System.SwitchPlayer
func Resolve[T data.Type](sys *System, action *action.Action) T {
	if !sys.isReady {
		panic(fmt.Errorf("input system is not initialized yet. " +
			"You must call SwitchPlayer at least once for one player configuration. " +
			"If you have multiplyer hot-seat, SwitchPlayer should be called every frame"))
	}

	// player not listen for this action currently
	// because another action map is active
	if action.Parent().ID() != sys.currentActionMapID {
		return *new(T)
	}

	// look for available controllers for this action
	ctls, ok := sys.binders[action.ID()]
	if !ok {
		return *new(T)
	}

	// look for binders (low-level input provider) for this action
	binders, ok := ctls[sys.currentControllerID]
	if !ok {
		return *new(T)
	}

	for _, binder := range binders {
		// binder is some controller, like [keyboard, gamepad, etc..]
		// when binder is called, it tries to resolve current value
		// with underlying device driver

		// resolvedValue must be some data.Type
		// if not, code should panic
		resolvedValue := binder()

		if typedValue, ok := resolvedValue.(T); ok {
			if data.IsHandled(typedValue) {
				return typedValue
			}

			// device is not handled by current controller:
			//   for example we can have this scheme [gamepad, keyboard]
			//   and if gamepad not connected to PC, value will be "Not handled"
			//   so we try process next to "keyboard" alternative controller
			continue
		}

		// not expected value
		var TType T
		panic(fmt.Errorf("resolved input value for action '%s.%s' not implement '%T' type. "+
			"Actual type-value: %T=`%v`",
			action.Parent().ID(),
			action.ID(),
			TType,
			resolvedValue,
			resolvedValue,
		))
	}

	// not have binders that can resolve this action
	// possible player not have attached controller
	// or all devices is lost (physically detached from pc)
	return *new(T)
}
