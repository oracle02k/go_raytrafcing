package hitable

import (
	"math"

	"github.com/oracle02k/go_raytracing/math3d"
	"github.com/oracle02k/go_raytracing/util3d"
)

type Sphere struct {
	center math3d.Vec3
	radius float64
}

func NewSphere(center math3d.Vec3, radius float64) *Sphere {
	return &Sphere{
		center: center,
		radius: radius,
	}
}

func (s *Sphere) Hit(r *util3d.Ray, t_min, t_max float64, rec *Record) bool {
	oc := r.Origin().Sub(s.center)
	a := math3d.Vec3Dot(r.Direction(), r.Direction())
	b := math3d.Vec3Dot(oc, r.Direction())
	c := math3d.Vec3Dot(oc, oc) - s.radius*s.radius
	discriminant := b*b - a*c

	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / a
		if temp < t_max && temp > t_min {
			rec.t = temp
			rec.p = r.PointAtParameter(rec.t)
			rec.normal = rec.p.Sub(s.center).Div(s.radius)
			return true
		}
		temp = (-b + math.Sqrt(discriminant)) / a
		if temp < t_max && temp > t_min {
			rec.t = temp
			rec.p = r.PointAtParameter(rec.t)
			rec.normal = rec.p.Sub(s.center).Div(s.radius)
			return true
		}
	}

	return false
}
