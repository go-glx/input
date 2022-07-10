package data

import "fmt"

type Type interface {
	any | Immediate | Stream | Vec1 | Vec2 | Vec3
}

func IsHandled(t Type) bool {
	switch value := t.(type) {
	case Immediate:
		return value.Handled
	case Stream:
		return value.Handled
	case Vec1:
		return value.Handled
	case Vec2:
		return value.Handled
	case Vec3:
		return value.Handled
	default:
		panic(fmt.Errorf("unknown data type '%T'", t))
	}
}

func NotHandledImmediate() Immediate {
	return Immediate{
		Handled:     false,
		JustPressed: false,
	}
}

func NotHandledStream() Stream {
	return Stream{
		Handled: false,
		Pressed: false,
	}
}

type Immediate struct {
	Handled     bool
	JustPressed bool
}

type Stream struct {
	Handled bool
	Pressed bool
}

type Vec1 struct {
	Handled bool
	V       float32
}

type Vec2 struct {
	Handled bool
	X       float32
	Y       float32
}

type Vec3 struct {
	Handled bool
	X       float32
	Y       float32
	Z       float32
}
