package util3d

import (
	"github.com/oracle02k/go_raytracing/math3d"
)

type Ray struct {
	origin    math3d.Vec3
	direction math3d.Vec3
}

func NewRay(origin, direction math3d.Vec3) *Ray {
	return &Ray{
		origin:    origin,
		direction: direction,
	}
}

func (r *Ray) Origin() math3d.Vec3 {
	return r.origin
}

func (r *Ray) Direction() math3d.Vec3 {
	return r.direction
}

func (r *Ray) PointAtParameter(t float64) math3d.Vec3 {
	return math3d.Vec3Add(r.origin, math3d.Vec3Scale(r.direction, t))
}
