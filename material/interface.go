package material

import (
	"github.com/oracle02k/go_raytracing/util3d"
	"github.com/oracle02k/go_raytracing/math3d"
	"github.com/oracle02k/go_raytracing/hit"
)

type Interface interface {
	Scatter(inR *util3d.Ray, rec *hit.Record) (bool, math3d.Vec3, *util3d.Ray)
}
