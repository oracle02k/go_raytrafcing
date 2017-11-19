package material

import (
	"github.com/oracle02k/go_raytracing/math3d"
	"github.com/oracle02k/go_raytracing/util3d"
	"math/rand"
	"github.com/oracle02k/go_raytracing/hit"
)

type Lambert struct {
	albedo math3d.Vec3
}

func NewLambert(albedo math3d.Vec3) *Lambert {
	return &Lambert{albedo: albedo}
}

func randomInUnitSphere() math3d.Vec3 {
	p := math3d.NewVec3(0, 0, 0)
	for {
		p = math3d.NewVec3(rand.Float64(), rand.Float64(), rand.Float64()).
			Scale(2).
			Sub(math3d.NewVec3(1, 1, 1))

		if p.SquaredLength() < 1.0 {
			break
		}
	}
	return p
}

func (l *Lambert) Scatter(inR *util3d.Ray, rec *hit.Record) (bool, math3d.Vec3, *util3d.Ray) {
	target := rec.P().Add(rec.Normal()).Add(randomInUnitSphere())
	scattered := util3d.NewRay(rec.P(), target.Sub(rec.P()))
	attenuation := l.albedo
	return true, attenuation, scattered
}
