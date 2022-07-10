package system

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-glx/input/bridge/keyboard/custom"
	"github.com/go-glx/input/system/controller"
	"github.com/go-glx/input/system/data"
	"github.com/go-glx/input/system/device"
	"github.com/go-glx/input/system/input/keyboard"
	"github.com/go-glx/input/system/player"
)

const (
	testActionMovement     = "movement"
	testActionMovementWalk = testActionMovement + "." + "walk"
	testActionMovementJump = testActionMovement + "." + "jump"

	testActionMenu       = "menu"
	testActionMenuSelect = testActionMenu + "." + "select"
)

func TestAssemble(t *testing.T) {
	// -- drivers
	testKeyboardDriver := custom.NewKeyboardDriver(true)

	// -- physical devices
	physicalKeyboard := keyboard.NewKeyboardDevice(testKeyboardDriver)

	// -- devices
	virtualKeyboardLeft := device.NewDevice("Keyboard Left", physicalKeyboard)
	virtualKeyboardRight := device.NewDevice("Keyboard Right", physicalKeyboard)

	// -- controllers
	controllerLeft := controller.NewController([]*device.Device{
		virtualKeyboardLeft,
	})
	controllerRight := controller.NewController([]*device.Device{
		virtualKeyboardRight,
	})

	// -- players
	player1 := player.NewPlayer()
	player1.AttachController(controllerLeft)
	// todo: player1.SwitchToActionMap( .. ) ??

	player2 := player.NewPlayer()
	player2.AttachController(controllerRight)

	// -- system
	system := NewSystem()

	// -- actions
	movementWalk := system.NewAction(testActionMovement, testActionMovementWalk)

	// -- bindings
	BindAs[data.Vec2](controllerLeft, movementWalk,
		system.FromVec2Keyboard(keyboard.KeyW, keyboard.KeyS, keyboard.KeyA, keyboard.KeyD), // WASD
	)
	BindAs[data.Vec2](controllerRight, movementWalk,
		system.FromVec2Keyboard(keyboard.KeyI, keyboard.KeyK, keyboard.KeyJ, keyboard.KeyL), // IJKL
	)

	// -- emulate test
	testKeyboardDriver.SetState(keyboard.KeyS, custom.StateDown) // pl1
	testKeyboardDriver.SetState(keyboard.KeyI, custom.StateDown) // pl2
	testKeyboardDriver.SetState(keyboard.KeyJ, custom.StateDown) // pl2

	// -- test
	system.SwitchPlayer(player1)
	walkVecPlr1 := Resolve[data.Vec2](system, testActionMovementWalk)
	system.SwitchPlayer(player2)
	walkVecPlr2 := Resolve[data.Vec2](system, testActionMovementWalk)

	assert.Equal(t, data.Vec2{Handled: true, X: 0, Y: 1}, walkVecPlr1)
	assert.Equal(t, data.Vec2{Handled: true, X: -1, Y: -1}, walkVecPlr2)
}
