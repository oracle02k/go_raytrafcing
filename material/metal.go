package material

import (
	"github.com/oracle02k/go_raytracing/util3d"
	"github.com/oracle02k/go_raytracing/math3d"
	"github.com/oracle02k/go_raytracing/hit"
)

type Metal struct {
	albedo math3d.Vec3
}

func NewMetal(albedo math3d.Vec3) *Metal {
	return &Metal{
		albedo: albedo,
	}
}

func (m *Metal) Scatter(inR *util3d.Ray, rec *hit.Record) (bool, math3d.Vec3, *util3d.Ray) {
	reflected := reflect(math3d.MakeUnitVector(inR.Direction()), rec.Normal())
	scattered := util3d.NewRay(rec.P(), reflected)
	attenuation := m.albedo
	return math3d.Vec3Dot(scattered.Direction(), rec.Normal()) > 0, attenuation, scattered
}

func reflect(v, n math3d.Vec3) math3d.Vec3 {
	return v.Sub(n.Scale(2.0).Scale(math3d.Vec3Dot(v, n)))
}
