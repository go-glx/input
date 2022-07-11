package raylib

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/go-glx/input/system/input/keyboard"
)

type KeyboardDriver struct {
}

func NewKeyboardDriver() *KeyboardDriver {
	return &KeyboardDriver{}
}

func (k *KeyboardDriver) IsEnabled() bool {
	return true
}

func (k *KeyboardDriver) OnEnabled(_ func())  {}
func (k *KeyboardDriver) OnDisabled(_ func()) {}

func (k *KeyboardDriver) IsJustPressed(key keyboard.Key) bool {
	if rlKey, supported := mapping[key]; supported {
		return rl.IsKeyPressed(rlKey)
	}

	return false
}

func (k *KeyboardDriver) IsJustReleased(key keyboard.Key) bool {
	if rlKey, supported := mapping[key]; supported {
		return rl.IsKeyReleased(rlKey)
	}

	return false
}

func (k *KeyboardDriver) IsDown(key keyboard.Key) bool {
	if rlKey, supported := mapping[key]; supported {
		return rl.IsKeyDown(rlKey)
	}

	return false
}
