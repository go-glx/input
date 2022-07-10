package system

import (
	"fmt"

	"github.com/go-glx/input/system/data"
)

// Resolve will fetch actual input values for provided action
// if you have more than 1 player, you need switch it with System.SwitchPlayer
func Resolve[T data.Type](action *Action) T {
	if !action.system.isReady {
		panic(fmt.Errorf("input system is not initialized yet. " +
			"You must call SwitchPlayer at least once for one player configuration. " +
			"If you have multiplyer hot-seat, SwitchPlayer should be called every frame"))
	}

	ctls, ok := action.system.binders[action.ID]
	if !ok {
		return *new(T)
	}

	binders, ok := ctls[action.system.currentControllerID]
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
			action.mapID,
			action.ID,
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
