package hitable

import (
	"github.com/oracle02k/go_raytracing/util3d"
)

type Interface interface {
	Hit(r *util3d.Ray, t_min, t_max float64, rec *Record) bool
}
