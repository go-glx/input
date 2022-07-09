package exp

type Resolver[T outputType] interface {
	IsAvailable() bool
	Resolve() T
}

type Layout[T outputType] struct {
	inputName any
	inputs    []Resolver[T]
}

// NewLayout will create Layout of type T for some specified input type
// example:
// 	type movement Vec2
// 	movementLayout := NewLayout[Vec2](
// 	  movement{},
// 	  InputVec2Keyboard(W,S,A,D)
// 	  InputVec2Keyboard(up,down,left,right)
// 	  InputVec2GamepadRightStick()
// 	)
func NewLayout[T outputType](inputName any, inputs ...Resolver[T]) *Layout[T] {
	return &Layout[T]{
		inputName: inputName,
		inputs:    inputs,
	}
}

func (lay Layout[T]) resolve() T {
	for _, input := range lay.inputs {
		if !input.IsAvailable() {
			continue
		}

		return input.Resolve()
	}

	return *new(T)
}
