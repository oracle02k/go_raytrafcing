package math3d

import (
	"math"
)

type Vec3 struct {
	e [3]float64
}

func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{e: [3]float64{x, y, z}}
}

func (v *Vec3) X() float64 {
	return v.e[0]
}

func (v *Vec3) Y() float64 {
	return v.e[1]
}

func (v *Vec3) Z() float64 {
	return v.e[2]
}
func (v *Vec3) R() float64 {
	return v.e[0]
}

func (v *Vec3) G() float64 {
	return v.e[1]
}

func (v *Vec3) B() float64 {
	return v.e[2]
}

func (v *Vec3) Length() float64 {
	return math.Sqrt(v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2])
}

func (v *Vec3) SquaredLength() float64 {
	return v.e[0]*v.e[0] + v.e[1]*v.e[1] + v.e[2]*v.e[2]
}

func (v *Vec3) MakeUnitVector() {
	k := 1.0 / math.Sqrt(v.e[0]*v.e[0]+v.e[1]*v.e[1]+v.e[2]*v.e[2])
	v.e[0] *= k
	v.e[1] *= k
	v.e[2] *= k
}

func Vec3Add(v1 *Vec3, v2 *Vec3) *Vec3 {
	return &Vec3{
		e: [3]float64{
			v1.e[0] + v2.e[0],
			v1.e[1] + v2.e[1],
			v1.e[2] + v2.e[2],
		},
	}
}

func Vec3Sub(v1 *Vec3, v2 *Vec3) *Vec3 {
	return &Vec3{
		e: [3]float64{
			v1.e[0] - v2.e[0],
			v1.e[1] - v2.e[1],
			v1.e[2] - v2.e[2],
		},
	}
}

func Vec3Mul(v1 *Vec3, v2 *Vec3) *Vec3 {
	return &Vec3{
		e: [3]float64{
			v1.e[0] * v2.e[0],
			v1.e[1] * v2.e[1],
			v1.e[2] * v2.e[2],
		},
	}
}

func Vec3Div(v1 *Vec3, v2 *Vec3) *Vec3 {
	return &Vec3{
		e: [3]float64{
			v1.e[0] / v2.e[0],
			v1.e[1] / v2.e[1],
			v1.e[2] / v2.e[2],
		},
	}
}

func Vec3Scale(v1 *Vec3, scale float64) *Vec3 {
	return &Vec3{
		e: [3]float64{
			v1.e[0] * scale,
			v1.e[1] * scale,
			v1.e[2] * scale,
		},
	}
}

func Vec3Dot(v1 *Vec3, v2 *Vec3) float64 {
	return v1.e[0]*v2.e[0] + v1.e[1]*v2.e[1] + v1.e[2]*v2.e[2]
}

func Vec3Closs(v1 *Vec3, v2 *Vec3) *Vec3 {
	return &Vec3{
		e: [3]float64{
			v1.e[1]*v2.e[2] - v1.e[2]*v2.e[1],
			-(v1.e[0]*v2.e[0] - v1.e[2]*v2.e[0]),
			v1.e[0]*v2.e[1] - v1.e[1]*v2.e[0],
		},
	}
}
