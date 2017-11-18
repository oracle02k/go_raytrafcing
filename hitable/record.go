package hitable

import "github.com/oracle02k/go_raytracing/math3d"

type Record struct {
	t      float64
	p      math3d.Vec3
	normal math3d.Vec3
}

func (h *Record) T() float64 {
	return h.t
}

func (h *Record) P() math3d.Vec3 {
	return h.p
}

func (h *Record) Normal() math3d.Vec3 {
	return h.normal
}
