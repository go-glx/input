package exp

import (
	input2 "github.com/go-glx/input/exp/input"
)

type KeyboardVec2 struct {
	up, down, left, right input2.KeyboardKey
	driver                *input2.VirtualKeyboard
}

func NewKeyboardVec2(
	up, down, left, right input2.KeyboardKey,
	driver *input2.VirtualKeyboard,
) *KeyboardVec2 {
	return &KeyboardVec2{
		up:     up,
		down:   down,
		left:   left,
		right:  right,
		driver: driver,
	}
}

func (k *KeyboardVec2) IsAvailable() bool {
	return k.driver.IsEnabled()
}

func (k *KeyboardVec2) Resolve() Vec2 {
	v := Vec2{}

	if k.driver.Driver().IsDown(k.up) {
		v.Y -= 1
	}
	if k.driver.Driver().IsDown(k.down) {
		v.Y += 1
	}
	if k.driver.Driver().IsDown(k.left) {
		v.X -= 1
	}
	if k.driver.Driver().IsDown(k.right) {
		v.X += 1
	}

	return v
}
