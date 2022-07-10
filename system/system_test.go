package system

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-glx/input/bridge/keyboard/custom"
	"github.com/go-glx/input/system/action"
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

func TestIntegrationFull(t *testing.T) {
	// -- actions
	movement := action.NewMap(testActionMovement)
	movementWalk := action.NewAction(movement, testActionMovementWalk)
	movementJump := action.NewAction(movement, testActionMovementJump)
	_ = movementJump // todo

	menu := action.NewMap(testActionMenu)
	menuSelect := action.NewAction(menu, testActionMenuSelect)

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
	player1.SwitchActionMap(movement)

	player2 := player.NewPlayer()
	player2.AttachController(controllerRight)
	player2.SwitchActionMap(movement)

	// -- system
	system := NewSystem()

	// -- bindings
	BindAs[data.Vec2](system, controllerLeft, movementWalk,
		system.FromVec2Keyboard(keyboard.KeyW, keyboard.KeyS, keyboard.KeyA, keyboard.KeyD), // WASD
	)
	BindAs[data.Vec1](system, controllerLeft, menuSelect,
		system.FromVec1Keyboard(keyboard.KeyW, keyboard.KeyS), // W(top) S(down)
	)
	BindAs[data.Vec2](system, controllerRight, movementWalk,
		system.FromVec2Keyboard(keyboard.KeyI, keyboard.KeyK, keyboard.KeyJ, keyboard.KeyL), // IJKL
	)
	BindAs[data.Vec1](system, controllerRight, menuSelect,
		system.FromVec1Keyboard(keyboard.Key9, keyboard.Key0), // 9(top) 0(down)
	)

	// -- emulate test
	testKeyboardDriver.SetState(keyboard.KeyS, custom.StateDown) // pl1
	testKeyboardDriver.SetState(keyboard.KeyI, custom.StateDown) // pl2
	testKeyboardDriver.SetState(keyboard.KeyJ, custom.StateDown) // pl2

	// -- test
	system.SwitchPlayer(player1)
	walkVecPlr1 := Resolve[data.Vec2](system, movementWalk)
	menuVecPlr1 := Resolve[data.Vec1](system, menuSelect)

	system.SwitchPlayer(player2)
	walkVecPlr2 := Resolve[data.Vec2](system, movementWalk)

	// -- check results
	assert.Equal(t, data.Vec2{Handled: true, X: 0, Y: 1}, walkVecPlr1)   // handled, movement map is active, keys is down
	assert.Equal(t, data.Vec1{Handled: false, V: 0}, menuVecPlr1)        // not handled, because menu map is not active
	assert.Equal(t, data.Vec2{Handled: true, X: -1, Y: -1}, walkVecPlr2) // player 2 handled as well
}
