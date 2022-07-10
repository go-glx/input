package system

import (
	"github.com/go-glx/input/system/data"
	"github.com/go-glx/input/system/input/keyboard"
	"github.com/go-glx/input/system/internal/binding"
)

// --------------------------------------
// --            IMMEDIATE             --
// --------------------------------------

// --------------------------------------
// --              STREAM              --
// --------------------------------------

func (s *System) FromStreamKeyboardKey(key keyboard.Key) binding.F[data.Stream] {
	return func() data.Stream {
		if s.currentKeyboard == nil {
			return data.NotHandledStream()
		}

		phys := s.currentKeyboard.Driver()
		return data.Stream{
			Handled: true,
			Pressed: phys.IsDown(key) || phys.IsJustPressed(key),
		}
	}
}

// --------------------------------------
// --               VEC1               --
// --------------------------------------

func (s *System) FromVec1(pos, neg binding.F[data.Stream]) binding.F[data.Vec1] {
	return binding.Vec1(pos, neg)
}

// --------------------------------------
// --               VEC2               --
// --------------------------------------

func (s *System) FromVec2(y, x binding.F[data.Vec1]) binding.F[data.Vec2] {
	return binding.Vec2(y, x)
}

func (s *System) FromVec2Keyboard(top, down, left, right keyboard.Key) binding.F[data.Vec2] {
	return s.FromVec2(
		s.FromVec1(
			s.FromStreamKeyboardKey(down),
			s.FromStreamKeyboardKey(top),
		),
		s.FromVec1(
			s.FromStreamKeyboardKey(right),
			s.FromStreamKeyboardKey(left),
		),
	)
}

// --------------------------------------
// --               VEC3               --
// --------------------------------------

func (s *System) FromVec3(y, x, z binding.F[data.Vec1]) binding.F[data.Vec3] {
	return binding.Vec3(y, x, z)
}
