package util3d

import (
	"github.com/oracle02k/go_raytracing/math3d"
	"math"
)

type Camera struct {
	origin          math3d.Vec3
	lowerLeftCorner math3d.Vec3
	horizontal      math3d.Vec3
	vertical        math3d.Vec3
}

func NewCamera(from, at, vup math3d.Vec3, vfov, aspect float64) *Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	w := math3d.MakeUnitVector(from.Sub(at))
	u := math3d.MakeUnitVector(math3d.Vec3Closs(vup, w))
	v := math3d.Vec3Closs(w, u)

	return &Camera{
		origin:          from,
		lowerLeftCorner: from.Sub(u.Scale(halfWidth)).Sub(v.Scale(halfHeight)).Sub(w),
		horizontal:      u.Scale(2 * halfWidth),
		vertical:        v.Scale(2 * halfHeight),
	}
}

func (c *Camera) Ray(u, v float64) *Ray {
	direction := c.lowerLeftCorner.
		Add(c.horizontal.Scale(u)).
		Add(c.vertical.Scale(v)).
		Sub(c.origin)

	return NewRay(c.origin, direction)
}
