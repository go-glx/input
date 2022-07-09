package custom

import (
	"github.com/go-glx/input/exp/input"
)

type State uint8

const (
	StateUp          State = 0
	StateJustPressed State = 1
	StateDown        State = 2
	StateReleased    State = 3
)

type KeyboardDriver struct {
	parentOnEnabled  func()
	parentOnDisabled func()
	enabled          bool
	keysState        map[input.KeyboardKey]State
}

func NewKeyboardDriver(enabled bool) *KeyboardDriver {
	return &KeyboardDriver{
		enabled:   enabled,
		keysState: make(map[input.KeyboardKey]State),
	}
}

func (kb *KeyboardDriver) SetState(key input.KeyboardKey, state State) {
	switch state {
	case StateUp:
		delete(kb.keysState, key)
		return
	default:
		kb.keysState[key] = state
	}
}

func (kb *KeyboardDriver) Enable() {
	if kb.enabled == true {
		return
	}

	kb.enabled = true
	if kb.parentOnEnabled != nil {
		kb.parentOnEnabled()
	}
}

func (kb *KeyboardDriver) Disable() {
	if kb.enabled == false {
		return
	}

	kb.enabled = false
	if kb.parentOnDisabled != nil {
		kb.parentOnDisabled()
	}
}

func (kb *KeyboardDriver) IsEnabled() bool {
	return kb.enabled
}

func (kb *KeyboardDriver) OnEnabled(fn func()) {
	kb.parentOnEnabled = fn
}

func (kb *KeyboardDriver) OnDisabled(fn func()) {
	kb.parentOnDisabled = fn
}

func (kb *KeyboardDriver) IsJustPressed(key input.KeyboardKey) bool {
	if state, exist := kb.keysState[key]; exist {
		return state == StateJustPressed
	}

	return false
}

func (kb *KeyboardDriver) IsJustReleased(key input.KeyboardKey) bool {
	if state, exist := kb.keysState[key]; exist {
		return state == StateReleased
	}

	return false
}

func (kb *KeyboardDriver) IsDown(key input.KeyboardKey) bool {
	if state, exist := kb.keysState[key]; exist {
		return state == StateDown
	}

	return false
}
