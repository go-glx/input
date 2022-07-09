package data

type Type interface {
	any | Flash | Vec1 | Vec2
}

type Flash struct {
	value bool
}

type Vec1 struct {
	X float32
}

type Vec2 struct {
	X float32
	Y float32
}
