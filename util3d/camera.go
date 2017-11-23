package util3d

import (
	"github.com/oracle02k/go_raytracing/math3d"
	"math"
	"math/rand"
)

type Camera struct {
	origin          math3d.Vec3
	lowerLeftCorner math3d.Vec3
	horizontal      math3d.Vec3
	vertical        math3d.Vec3
	u               math3d.Vec3
	v               math3d.Vec3
	w               math3d.Vec3
	lensRadius      float64
}

func NewCamera(from, at, vup math3d.Vec3, vfov, aspect, aperture, focusDist float64) *Camera {
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	w := math3d.MakeUnitVector(from.Sub(at))
	u := math3d.MakeUnitVector(math3d.Vec3Closs(vup, w))
	v := math3d.Vec3Closs(w, u)

	return &Camera{
		origin:          from,
		lowerLeftCorner: from.Sub(u.Scale(halfWidth * focusDist)).Sub(v.Scale(halfHeight * focusDist)).Sub(w.Scale(focusDist)),
		horizontal:      u.Scale(2 * halfWidth * focusDist),
		vertical:        v.Scale(2 * halfHeight * focusDist),
		u:               u,
		v:               v,
		w:               w,
		lensRadius:      aperture / 2,
	}
}

func (c *Camera) Ray(u, v float64) *Ray {

	rd := randomInUnitDesk().Scale(c.lensRadius)
	offset := c.u.Scale(rd.X()).Add(c.v.Scale(rd.Y()))

	direction := c.lowerLeftCorner.
		Add(c.horizontal.Scale(u)).
		Add(c.vertical.Scale(v)).
		Sub(c.origin).
		Sub(offset)

	return NewRay(c.origin.Add(offset), direction)
}

func randomInUnitDesk() math3d.Vec3 {
	p := math3d.NewVec3(0, 0, 0)
	for {
		p = math3d.NewVec3(rand.Float64(), rand.Float64(), 0.0).Scale(2).Sub(math3d.NewVec3(1, 1, 0))
		if math3d.Vec3Dot(p, p) < 1.0 {
			break
		}
	}
	return p
}
