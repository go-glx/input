package exp

type outputType interface {
	Vec1 | Vec2 | Action | Mouse
}

// Vec1 is single value bound to [-1 .. 0 .. +1]
type Vec1 struct {
	X float32
}

// Vec2 is two values bound to [-1 .. 0 .. +1]
// Usually is every kind of movement inputs (WSAD, arrows, stick, etc..)
type Vec2 struct {
	X float32
	Y float32
}

// Action is one time event that just happened
// in last frame, is good for:
//  - button just pressed
//  - mouse just clicked
//  - etc..
type Action struct {
	IsHappened bool
}

// Mouse used only for define current mouse coords
type Mouse struct {
	X int
	Y int
}
