package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/oracle02k/go_raytracing/math3d"
	"github.com/oracle02k/go_raytracing/util3d"
)

func hitSphere(center math3d.Vec3, radius float64, r *util3d.Ray) float64 {
	oc := r.Origin().Sub(center)
	a := math3d.Vec3Dot(r.Direction(), r.Direction())
	b := 2.0 * math3d.Vec3Dot(oc, r.Direction())
	c := math3d.Vec3Dot(oc, oc) - radius*radius
	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return -1.0
	}

	return (-b - math.Sqrt(discriminant)) / (2.0 * a)
}

func color(r *util3d.Ray) math3d.Vec3 {
	t := hitSphere(math3d.NewVec3(0, 0, -1), 0.5, r)
	if t > 0.0 {
		n := math3d.MakeUnitVector(r.PointAtParameter(t).Sub(math3d.NewVec3(0, 0, -1)))
		return math3d.NewVec3(n.X()+1, n.Y()+1, n.Z()+1).Scale(0.5)
	}

	unitDirection := math3d.MakeUnitVector(r.Direction())
	t = 0.5 * (unitDirection.Y() + 1.0)
	return math3d.NewVec3(1.0, 1.0, 1.0).Scale(1.0 - t).Add(math3d.NewVec3(0.5, 0.7, 1.0).Scale(t))
}

func main() {

	nx := 200
	ny := 100

	writeFile, _ := os.OpenFile("test.ppm", os.O_WRONLY|os.O_CREATE, 0600)
	writer := bufio.NewWriter(writeFile)
	writer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))

	lowerLeftCenter := math3d.NewVec3(-2.0, -1.0, -1.0)
	horizontal := math3d.NewVec3(4.0, 0.0, 0.0)
	vertical := math3d.NewVec3(0.0, 2.0, 0.0)
	origin := math3d.NewVec3(0.0, 0.0, 0.0)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			direction := lowerLeftCenter.Add(math3d.Vec3Scale(horizontal, u)).Add(math3d.Vec3Scale(vertical, v))
			r := util3d.NewRay(origin, direction)
			col := color(r).Scale(255.99)

			ir := int32(col.R())
			ig := int32(col.G())
			ib := int32(col.B())
			writer.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}
	writer.Flush()
	writeFile.Close()
}
