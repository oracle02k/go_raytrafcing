package material

import (
	"github.com/oracle02k/go_raytracing/math3d"
	"math"
	"github.com/oracle02k/go_raytracing/util3d"
	"github.com/oracle02k/go_raytracing/hit"
)

type Dielectric struct {
	refIdx float64
}

func NewDielectric(ri float64) *Dielectric {
	return &Dielectric{
		refIdx: ri,
	}
}

func (d *Dielectric) Scatter(inR *util3d.Ray, rec *hit.Record) (bool, math3d.Vec3, *util3d.Ray) {

	attenuation := math3d.NewVec3(1.0, 1.0, 1.0)
	var outWardNormal math3d.Vec3
	var niOverNt float64

	if math3d.Vec3Dot(inR.Direction(), rec.Normal()) > 0 {
		outWardNormal = rec.Normal().Minus()
		niOverNt = d.refIdx
	} else {
		outWardNormal = rec.Normal()
		niOverNt = 1.0 / d.refIdx
	}

	result, refracted := refract(inR.Direction(), outWardNormal, niOverNt)
	if result {
		scattered := util3d.NewRay(rec.P(), refracted);
		return true, attenuation, scattered
	}

	reflected := reflect(inR.Direction(), rec.Normal())
	scattered := util3d.NewRay(rec.P(), reflected);
	return false, attenuation, scattered
}

func refract(v, n math3d.Vec3, niOverNt float64) (bool, math3d.Vec3) {
	uv := math3d.MakeUnitVector(v)
	dt := math3d.Vec3Dot(uv, n)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		refracted := uv.Sub(n.Scale(dt)).Scale(niOverNt).Sub(n.Scale(math.Sqrt(discriminant)))
		return true, refracted
	}
	return false, math3d.NewVec3(0, 0, 0)
}
