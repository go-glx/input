package system

import (
	"fmt"

	"github.com/go-glx/input/system/controller"
	"github.com/go-glx/input/system/device"
	"github.com/go-glx/input/system/input/keyboard"
	"github.com/go-glx/input/system/internal/binding"
	"github.com/go-glx/input/system/player"
)

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

// NewAction will create Action that can be bind to system
// later wia BindAs function
func (s *System) NewAction(mapID string, id string) *Action {
	return newAction(s, actionID(id), actionMapID(mapID))
}

// SwitchPlayer will set current player for who lib should
// resolve provided input actions wia Resolve.
//
// If you have more than one player, you MUST call this
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
