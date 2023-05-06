package sdl2

import "github.com/go-glx/input/system/input/keyboard"

const (
	keyboardStatePressed keyboardState = iota
	keyboardStateReleased
	keyboardStateDown
)

type (
	keyboardState uint8

	queuedAction struct {
		action keyboardState
		key    keyboard.Key
	}
)

type state struct {
	keys   map[keyboard.Key]keyboardState
	queued []queuedAction
}

func newState() *state {
	return &state{
		keys:   make(map[keyboard.Key]keyboardState),
		queued: make([]queuedAction, 0, 32),
	}
}

func (s *state) onTick() {
	s.clean()

	for _, change := range s.queued {
		s.keys[change.key] = change.action
	}

	s.queued = make([]queuedAction, 0, 32)
}

func (s *state) queue(key keyboard.Key, action keyboardState) {
	s.queued = append(s.queued, queuedAction{
		action: action,
		key:    key,
	})
}

func (s *state) clean() {
	// pressed -> down
	for key, state := range s.keys {
		if state == keyboardStatePressed {
			s.keys[key] = keyboardStateDown
		}
	}

	// released -> up (up = default = delete)
	for key, state := range s.keys {
		if state == keyboardStateReleased {
			delete(s.keys, key)
		}
	}
}

func (s *state) isPressed(key keyboard.Key) bool {
	if state, exist := s.keys[key]; exist {
		return state == keyboardStatePressed
	}

	return false
}

func (s *state) isReleased(key keyboard.Key) bool {
	if state, exist := s.keys[key]; exist {
		return state == keyboardStateReleased
	}

	return false
}

func (s *state) isDown(key keyboard.Key) bool {
	if state, exist := s.keys[key]; exist {
		return state == keyboardStateDown
	}

	return false
}
