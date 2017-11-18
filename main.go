package main

import (
	"bufio"
	"fmt"
	"math"
	"os"

	"github.com/oracle02k/go_raytracing/math3d"
	"github.com/oracle02k/go_raytracing/util3d"
	"github.com/oracle02k/go_raytracing/hitable"
)

func color(r *util3d.Ray, world *hitable.List) math3d.Vec3 {
	rec := &hitable.Record{}
	if world.Hit(r, 0.0, math.MaxFloat64, rec) {
		normal := rec.Normal()
		return math3d.NewVec3(normal.X()+1.0, normal.Y()+1.0, normal.Z()+1.0).Scale(0.5)
	}

	unit := math3d.MakeUnitVector(r.Direction())
	t := 0.5 * (unit.Y() + 1.0)
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

	world := hitable.NewList()
	world.AddHitable(hitable.NewSphere(math3d.NewVec3(0, 0, -1), 0.5))
	world.AddHitable(hitable.NewSphere(math3d.NewVec3(0, -100.5, -1), 100))

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float64(i) / float64(nx)
			v := float64(j) / float64(ny)

			direction := lowerLeftCenter.Add(math3d.Vec3Scale(horizontal, u)).Add(math3d.Vec3Scale(vertical, v))
			r := util3d.NewRay(origin, direction)
			col := color(r, world).Scale(255.99)

			ir := int32(col.R())
			ig := int32(col.G())
			ib := int32(col.B())
			writer.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}
	writer.Flush()
	writeFile.Close()
}
