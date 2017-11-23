package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"math/rand"

	"github.com/oracle02k/go_raytracing/hitable"
	"github.com/oracle02k/go_raytracing/math3d"
	"github.com/oracle02k/go_raytracing/util3d"
	"github.com/oracle02k/go_raytracing/material"
)

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

func color(r *util3d.Ray, world *hitable.List, depth int32) math3d.Vec3 {
	hit, hitable, record := world.Hit(r, 0.001, math.MaxFloat64)
	if hit {
		if depth < 50 {
			result, attenuation, scattered := hitable.Material().Scatter(r, record)
			if result {
				return math3d.Vec3Mul(attenuation, color(scattered, world, depth+1))
			}
			return math3d.NewVec3(0, 0, 0);
		}
		return math3d.NewVec3(0, 0, 0);
	}

	unit := math3d.MakeUnitVector(r.Direction())
	t := 0.5 * (unit.Y() + 1.0)
	return math3d.NewVec3(1.0, 1.0, 1.0).Scale(1.0 - t).Add(math3d.NewVec3(0.5, 0.7, 1.0).Scale(t))
}

func main() {

	nx := 400
	ny := 200
	ns := 100

	writeFile, _ := os.OpenFile("test.ppm", os.O_WRONLY|os.O_CREATE, 0600)
	writer := bufio.NewWriter(writeFile)
	writer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))

	world := hitable.NewList()
	world.AddHitable(hitable.NewSphere(math3d.NewVec3(0, 0, -1), 0.5, material.NewLambert(math3d.NewVec3(0.8, 0.3, 0.3))))
	world.AddHitable(hitable.NewSphere(math3d.NewVec3(0, -100.5, -1), 100, material.NewLambert(math3d.NewVec3(0.8, 0.8, 0.0))))
	world.AddHitable(hitable.NewSphere(math3d.NewVec3(1, 0, -1), 0.5, material.NewMetal(math3d.NewVec3(0.8, 0.6, 0.2), 1.0)))
	world.AddHitable(hitable.NewSphere(math3d.NewVec3(-1, 0, -1), 0.5, material.NewDielectric(1.5)))

	camera := util3d.NewCamera(
		math3d.NewVec3(0, 1, 1),
		math3d.NewVec3(0, 0, -1),
		math3d.NewVec3(0, 1, 0),
		30,
		float64(nx)/float64(ny),
	)

	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			col := math3d.NewVec3(0, 0, 0)
			for s := 0; s < ns; s++ {
				u := (float64(i) + rand.Float64()) / float64(nx);
				v := (float64(j) + rand.Float64()) / float64(ny);
				r := camera.Ray(u, v)
				col = col.Add(color(r, world, 0));
			}

			col = col.Div(float64(ns))
			col = math3d.NewVec3(math.Sqrt(col.X()), math.Sqrt(col.Y()), math.Sqrt(col.Z())).Scale(255.99)
			ir := int32(col.R())
			ig := int32(col.G())
			ib := int32(col.B())
			writer.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}

	writer.Flush()
	writeFile.Close()
}
