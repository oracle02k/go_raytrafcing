package util3d

import (
	"github.com/oracle02k/go_raytracing/math3d"
)

type Camera struct {
	origin          math3d.Vec3
	lowerLeftCorner math3d.Vec3
	horizontal      math3d.Vec3
	vertical        math3d.Vec3
}

func NewCamera() *Camera {
	return &Camera{
		origin:          math3d.NewVec3(0.0, 0.0, 0.0),
		lowerLeftCorner: math3d.NewVec3(-2.0, -1.0, -1.0),
		horizontal:      math3d.NewVec3(4.0, 0.0, 0.0),
		vertical:        math3d.NewVec3(0.0, 2.0, 0.0),
	}
}

func (c *Camera) Ray(u, v float64) *Ray {
	direction := c.lowerLeftCorner.
		Add(c.horizontal.Scale(u)).
		Add(c.vertical.Scale(v)).
		Sub(c.origin)

	return NewRay(c.origin, direction)
}
