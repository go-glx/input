package exp

import (
	input2 "github.com/go-glx/input/exp/input"
)

type KeyboardVec1 struct {
	left, right input2.KeyboardKey
	driver      *input2.VirtualKeyboard
}

func NewKeyboardVec1(
	left, right input2.KeyboardKey,
	driver *input2.VirtualKeyboard,
) *KeyboardVec2 {
	return &KeyboardVec2{
		left:   left,
		right:  right,
		driver: driver,
	}
}

func (k *KeyboardVec1) IsAvailable() bool {
	return k.driver.IsEnabled()
}

func (k *KeyboardVec1) Resolve() Vec1 {
	v := Vec1{}

	if k.driver.Driver().IsDown(k.left) {
		v.X -= 1
	}
	if k.driver.Driver().IsDown(k.right) {
		v.X += 1
	}

	return v
}
