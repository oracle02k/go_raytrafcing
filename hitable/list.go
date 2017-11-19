package hitable

import (
	"github.com/oracle02k/go_raytracing/util3d"
	"github.com/oracle02k/go_raytracing/hit"
)

type List struct {
	hitables []Interface
}

func NewList() *List {
	return &List{hitables: make([]Interface, 0, 32)}
}

func (l *List) AddHitable(hitable Interface) {
	l.hitables = append(l.hitables, hitable)
}

func (l *List) Hit(r *util3d.Ray, t_min, t_max float64) (bool, Interface, *hit.Record) {
	var hitable Interface
	var record *hit.Record
	hitAnything  := false

	closetSoFar := t_max
	for i := 0; i < len(l.hitables); i++ {
		result, r := l.hitables[i].Hit(r, t_min, closetSoFar)
		if result {
			record = r
			closetSoFar = r.T()
			hitable = l.hitables[i]
			hitAnything = true
		}
	}

	return hitAnything, hitable, record
}
