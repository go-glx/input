package sdl2

import (
	"github.com/go-glx/input/system/input/keyboard"
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardDriver struct {
	state *state
}

func NewKeyboardDriver() *KeyboardDriver {
	return &KeyboardDriver{
		state: newState(),
	}
}

func (k *KeyboardDriver) HandleSDLKeyboardEvent(event *sdl.KeyboardEvent) {
	key, exist := mapping[event.Keysym.Sym]
	if !exist {
		return
	}

	action, exist := actionMapping[event.State]
	if !exist {
		return
	}

	if !(action == keyboardStatePressed || action == keyboardStateReleased) {
		return
	}

	k.state.queue(key, action)
}

func (k *KeyboardDriver) IsEnabled() bool {
	return true
}

func (k *KeyboardDriver) Tick() {
	k.state.onTick()
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
