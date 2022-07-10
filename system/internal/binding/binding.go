package binding

import "github.com/go-glx/input/system/data"

type (
	F[T data.Type] func() T
)

func Vec1(pos, neg F[data.Stream]) F[data.Vec1] {
	return func() data.Vec1 {
		v, handled := normalizeStream(pos, neg)

		return data.Vec1{
			V:       v,
			Handled: handled,
		}
	}
}

func Vec2(y, x F[data.Vec1]) F[data.Vec2] {
	return func() data.Vec2 {
		xx, yy := x(), y()
		return data.Vec2{
			Y:       yy.V,
			X:       xx.V,
			Handled: xx.Handled && yy.Handled,
		}
	}
}

func Vec3(x, y, z F[data.Vec1]) F[data.Vec3] {
	return func() data.Vec3 {
		xx, yy, zz := x(), y(), z()
		return data.Vec3{
			X:       xx.V,
			Y:       yy.V,
			Z:       zz.V,
			Handled: xx.Handled && yy.Handled && zz.Handled,
		}
	}
}

func normalizeStream(pos, neg F[data.Stream]) (float32, bool) {
	posV, negV := pos(), neg()

	v := float32(0.0)

	if negV.Pressed {
		v -= 1
	}
	if posV.Pressed {
		v += 1
	}

	return v, posV.Handled && negV.Handled
}
