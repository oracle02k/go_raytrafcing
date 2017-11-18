package math3d

import (
	"math"
)

type Vec3 struct {
	e [3]float64
}

func NewVec3(x, y, z float64) Vec3 {
	return Vec3{e: [3]float64{x, y, z}}
}

func (v Vec3) X() float64 {
	return v.e[0]
}

func (v Vec3) Y() float64 {
	return v.e[1]
}

func (v Vec3) Z() float64 {
	return v.e[2]
}
func (v Vec3) R() float64 {
	return v.e[0]
}

func (v Vec3) G() float64 {
	return v.e[1]
}

func (v Vec3) B() float64 {
	return v.e[2]
}

func (v Vec3) Length() float64 {
	return math.Sqrt(v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2])
}

func (v Vec3) SquaredLength() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v Vec3) Add(add Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			v.e[0] + add.e[0],
			v.e[1] + add.e[1],
			v.e[2] + add.e[2],
		},
	}
}

func (v Vec3) Sub(sub Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			v.e[0] - sub.e[0],
			v.e[1] - sub.e[1],
			v.e[2] - sub.e[2],
		},
	}
}

func (v Vec3) Scale(scale float64) Vec3 {
	return Vec3{
		e: [3]float64{
			v.e[0] * scale,
			v.e[1] * scale,
			v.e[2] * scale,
		},
	}
}

func (v Vec3) Div(div float64) Vec3 {
	return Vec3{
		e: [3]float64{
			v.e[0] / div,
			v.e[1] / div,
			v.e[2] / div,
		},
	}
}

func MakeUnitVector(v Vec3) Vec3 {
	l := math.Sqrt(v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2])
	k := 1.0 / l
	v.e[0] *= k
	v.e[1] *= k
	v.e[2] *= k

	return v
}

func Vec3Add(v1, v2 Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[0] + v2.e[0],
			v1.e[1] + v2.e[1],
			v1.e[2] + v2.e[2],
		},
	}
}

func Vec3Sub(v1, v2 Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[0] - v2.e[0],
			v1.e[1] - v2.e[1],
			v1.e[2] - v2.e[2],
		},
	}
}

func Vec3Mul(v1, v2 Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[0] * v2.e[0],
			v1.e[1] * v2.e[1],
			v1.e[2] * v2.e[2],
		},
	}
}

func Vec3Div(v1, v2 Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[0] / v2.e[0],
			v1.e[1] / v2.e[1],
			v1.e[2] / v2.e[2],
		},
	}
}

func Vec3Scale(v1 Vec3, scale float64) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[0] * scale,
			v1.e[1] * scale,
			v1.e[2] * scale,
		},
	}
}

func Vec3Dot(v1 Vec3, v2 Vec3) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}

func Vec3Closs(v1 Vec3, v2 Vec3) Vec3 {
	return Vec3{
		e: [3]float64{
			v1.e[1]*v2.e[2] - v1.e[2]*v2.e[1],
			-(v1.e[0]*v2.e[0] - v1.e[2]*v2.e[0]),
			v1.e[0]*v2.e[1] - v1.e[1]*v2.e[0],
		},
	}
}
