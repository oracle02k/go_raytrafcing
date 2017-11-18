package util3d

import (
	"github.com/oracle02k/go_raytracing/math3d"
)

type HitRecord struct {
	t      float64
	p      math3d.Vec3
	normal math3d.Vec3
}

func (h *HitRecord) T() float64 {
	return h.t
}

func (h *HitRecord) P() math3d.Vec3 {
	return h.p
}

func (h *HitRecord) Normal() math3d.Vec3 {
	return h.normal
}

type Hitable interface {
	Hit(r *Ray, t_min, t_max float64, rec *HitRecord) bool
}
