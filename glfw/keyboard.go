package glfw

import (
	"github.com/go-gl/glfw/v3.3/glfw"

	"github.com/go-glx/input/system/input/keyboard"
)

type KeyboardDriver struct {
	state *state
}

func NewKeyboardDriver(window *glfw.Window) *KeyboardDriver {
	drv := &KeyboardDriver{
		state: newState(),
	}

	window.SetKeyCallback(drv.onInput)

	return drv
}

func (k *KeyboardDriver) onInput(_ *glfw.Window, glfwKey glfw.Key, _ int, glfwAction glfw.Action, _ glfw.ModifierKey) {
	key, exist := mapping[glfwKey]
	if !exist {
		return
	}

	action, exist := actionMapping[glfwAction]
	if !exist {
		return
	}

	if !(action == keyboardStatePressed || action == keyboardStateReleased) {
		return
	}

	k.state.queue(key, action)
}

func (k *KeyboardDriver) Tick() {
	k.state.onTick()
}

func (k *KeyboardDriver) IsEnabled() bool {
	return true
}

func (k *KeyboardDriver) OnEnabled(_ func())  {}
func (k *KeyboardDriver) OnDisabled(_ func()) {}

func (k *KeyboardDriver) IsJustPressed(key keyboard.Key) bool {
	return k.state.isPressed(key)
}

func (k *KeyboardDriver) IsJustReleased(key keyboard.Key) bool {
	return k.state.isReleased(key)
}

func (k *KeyboardDriver) IsDown(key keyboard.Key) bool {
	return k.state.isDown(key)
}
