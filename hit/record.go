package hit

import (
	"github.com/oracle02k/go_raytracing/math3d"
)

type Record struct {
	t      float64
	p      math3d.Vec3
	normal math3d.Vec3
}

func NewRecord(t float64, p, normal math3d.Vec3) *Record {
	return &Record{
		t:      t,
		p:      p,
		normal: normal,
	}
}

func (r *Record) T() float64 {
	return r.t
}

func (r *Record) P() math3d.Vec3 {
	return r.p
}

func (r *Record) Normal() math3d.Vec3 {
	return r.normal
}
