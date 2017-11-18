package hitable

import "github.com/oracle02k/go_raytracing/util3d"

type List struct {
	hitables []Interface
}

func NewList() *List {
	return &List{hitables: make([]Interface, 0, 32)}
}

func (l *List) AddHitable(hitable Interface) {
	l.hitables = append(l.hitables, hitable)
}

func (l *List) Hit(r *util3d.Ray, t_min, t_max float64, rec *Record) bool {
	tmpRec := Record{}
	hitAnything := false
	closetSoFar := t_max
	for i := 0; i < len(l.hitables); i++ {
		if l.hitables[i].Hit(r, t_min, closetSoFar, rec) {
			hitAnything = true
			closetSoFar = tmpRec.t
			rec = &tmpRec
		}
	}
	return hitAnything
}
