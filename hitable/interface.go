package hitable

import (
	"github.com/oracle02k/go_raytracing/util3d"
	"github.com/oracle02k/go_raytracing/material"
	"github.com/oracle02k/go_raytracing/hit"
)

type Interface interface {
	Hit(r *util3d.Ray, t_min, t_max float64) (bool, Interface, *hit.Record)
	Material() material.Interface
}
