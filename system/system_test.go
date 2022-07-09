package system

import (
	"testing"

	"github.com/go-glx/input/system/binding"
	"github.com/go-glx/input/system/data"
	"github.com/go-glx/input/system/device"
)

type testActionIdJump struct{}
type testActionIdWalking struct{}
type testActionIdSelectMenu struct{}

func TestNewSystem(t *testing.T) {
	// -- Devices
	// both keyboards can reference ONE physical keyboard
	keyboard1 := device.NewDevice(device.KindKeyboard, 1, "Keyboard 1")
	keyboard2 := device.NewDevice(device.KindKeyboard, 2, "Keyboard 2")

	// both gamepad can reference DIFFERENT physical gamepads
	gamepad1 := device.NewDevice(device.KindGamepad, 1, "Gamepad 1")
	gamepad2 := device.NewDevice(device.KindGamepad, 2, "Gamepad 2")

	// -- Players
	playerOne := NewPlayer(Player1, true)
	playerOne.AttachDevice(keyboard1)
	playerOne.AttachDevice(gamepad1)

	playerTwo := NewPlayer(Player2, true)
	playerTwo.AttachDevice(keyboard2)
	playerTwo.AttachDevice(gamepad2)

	players := [4]*Player{
		playerOne,
		playerTwo,
		NewPlayer(Player3, false),
		NewPlayer(Player4, false),
	}

	// -- Action maps

	actionMapMovement := NewActionMap("movement")

	BindTo[data.Flash](actionMapMovement, testActionIdJump{}, []*binding.Binding{
		binding.NewBinding(keyboard1, "LCtrl"),
		binding.NewBinding(keyboard2, "RCtrl"),
		binding.NewBinding(gamepad1, "X"),
		binding.NewBinding(gamepad2, "X"),
	})
	BindTo[data.Vec2](actionMapMovement, testActionIdWalking{}, []*binding.Binding{
		binding.NewBinding(keyboard1, "WASD"),
		binding.NewBinding(keyboard2, "ARROWS"),
		binding.NewBinding(gamepad1, "RightStick"),
		binding.NewBinding(gamepad2, "RightStick"),
	})

	actionMapMenu := NewActionMap("menu")

	BindTo[data.Vec1](actionMapMovement, testActionIdSelectMenu{}, []*binding.Binding{
		binding.NewBinding(keyboard1, "WS"),
		binding.NewBinding(keyboard2, "Up,Down"),
		binding.NewBinding(gamepad1, "RightStickUp"),
		binding.NewBinding(gamepad2, "RightStickDown"),
	})

	actionMaps := []*ActionMap{
		actionMapMovement,
		actionMapMenu,
	}

	// -- System

	system := NewSystem(players, actionMaps)

	// -- Tests
	_ = system
}
