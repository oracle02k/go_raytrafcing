package hitable

import (
	"math"

	"github.com/oracle02k/go_raytracing/math3d"
	"github.com/oracle02k/go_raytracing/util3d"
	"github.com/oracle02k/go_raytracing/material"
	"github.com/oracle02k/go_raytracing/hit"
)

type Sphere struct {
	center   math3d.Vec3
	radius   float64
	material material.Interface
}

func NewSphere(center math3d.Vec3, radius float64, material material.Interface) *Sphere {
	return &Sphere{
		center:   center,
		radius:   radius,
		material: material,
	}
}

func (s *Sphere) Hit(r *util3d.Ray, t_min, t_max float64) (bool, *hit.Record) {
	oc := r.Origin().Sub(s.center)
	a := math3d.Vec3Dot(r.Direction(), r.Direction())
	b := math3d.Vec3Dot(oc, r.Direction())
	c := math3d.Vec3Dot(oc, oc) - s.radius*s.radius
	discriminant := b*b - a*c

	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / a
		if temp < t_max && temp > t_min {
			p := r.PointAtParameter(temp)
			rec := hit.NewRecord(temp, p, p.Sub(s.center).Div(s.radius))
			return true, rec
		}
		temp = (-b + math.Sqrt(discriminant)) / a
		if temp < t_max && temp > t_min {
			p := r.PointAtParameter(temp)
			rec := hit.NewRecord(temp, p, p.Sub(s.center).Div(s.radius))
			return true, rec
		}
	}

	return false, nil
}

func (s *Sphere) Material() material.Interface {
	return s.material
}
