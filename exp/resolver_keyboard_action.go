package exp

import (
	input2 "github.com/go-glx/input/exp/input"
)

type keyboardActionType = uint8

const (
	keyboardOnJustPressed  keyboardActionType = 0
	keyboardOnJustReleased keyboardActionType = 0
)

type KeyboardOn struct {
	actionType keyboardActionType
	key        input2.KeyboardKey
	driver     *input2.VirtualKeyboard
}

func NewKeyboardOnPressed(
	key input2.KeyboardKey,
	driver *input2.VirtualKeyboard,
) *KeyboardOn {
	return &KeyboardOn{
		actionType: keyboardOnJustPressed,
		key:        key,
		driver:     driver,
	}
}

func NewKeyboardOnReleased(
	key input2.KeyboardKey,
	driver *input2.VirtualKeyboard,
) *KeyboardOn {
	return &KeyboardOn{
		actionType: keyboardOnJustReleased,
		key:        key,
		driver:     driver,
	}
}

func (k *KeyboardOn) IsAvailable() bool {
	return k.driver.IsEnabled()
}

func (k *KeyboardOn) Resolve() Action {
	if k.actionType == keyboardOnJustPressed && k.driver.Driver().IsJustPressed(k.key) {
		return Action{IsHappened: true}
	}

	if k.actionType == keyboardOnJustReleased && k.driver.Driver().IsJustReleased(k.key) {
		return Action{IsHappened: true}
	}

	return Action{IsHappened: false}
}
