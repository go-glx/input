package system

import (
	"fmt"

	"github.com/go-glx/input/system/controller"
	"github.com/go-glx/input/system/data"
	"github.com/go-glx/input/system/device"
	"github.com/go-glx/input/system/input/keyboard"
	"github.com/go-glx/input/system/internal/binding"
	"github.com/go-glx/input/system/player"
)

type actionID string

type System struct {
	isReady bool

	currentControllerID controller.ID
	currentKeyboard     *keyboard.PhysicalDevice
	// currentGamepad *VirtualGamepad // todo
	// currentMouse *VirtualMouse  // todo

	binders map[actionID]map[controller.ID][]binding.F[any]
}

func NewSystem() *System {
	return &System{
		isReady: false,
	}
}

// Resolve will fetch actual input values for provided action
// if you have more than 1 player, you need switch it with System.SwitchPlayer
func Resolve[T data.Type](s *System, action actionID) T {
	if !s.isReady {
		panic(fmt.Errorf("input system is not initialized yet. " +
			"You must call SwitchPlayer at least once for one player configuration. " +
			"If you have multiplyer hot-seat, SwitchPlayer should be called every frame"))
	}

	ctls, ok := s.binders[action]
	if !ok {
		return *new(T)
	}

	binders, ok := ctls[s.currentControllerID]
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
		panic(fmt.Errorf("resolved input value for action '%s' not implement '%T' type. "+
			"Actual type-value: %T=`%v`",
			action,
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

// SwitchPlayer will set current player for who need
// resolve all input keys.
// If you have more than 1 player, you must call this
// function every frame, right before Resolve
//
// example:
// 	SwitchPlayer(1)
//	plr1Moving := Resolve(actionMove)
//	plr1Jumping := Resolve(actionJump)
// 	SwitchPlayer(2)
//	plr2Moving := Resolve(actionMove)
//	plr2Jumping := Resolve(actionJump)
//
// In one player configuration you must call this at least once
//
func (s *System) SwitchPlayer(plr *player.Player) {
	if !s.isReady {
		s.isReady = true
	}

	// reset current player controller
	s.reset()

	// check if this player has controller
	ctl := plr.Controller()
	if ctl == nil {
		return
	}

	// set new player controller
	s.setCurrentController(ctl)
}

func (s *System) reset() {
	s.currentControllerID = ""
	s.currentKeyboard = nil
	// todo: gamepad
	// todo: mouse
}

func (s *System) setCurrentController(ctl *controller.Controller) {
	s.currentControllerID = ctl.ID()

	for _, logicalDevice := range ctl.Devices() {
		s.setCurrentDevice(logicalDevice)
	}
}

func (s *System) setCurrentDevice(virtual *device.Device) {
	switch physical := virtual.Physical().(type) {
	// todo: gamepad, mouse
	case *keyboard.PhysicalDevice:
		s.currentKeyboard = physical
		return
	default:
		panic(fmt.Errorf("unsupported device '%s' of type '%s'",
			virtual.Name(),
			virtual.Kind(),
		))
	}
}
