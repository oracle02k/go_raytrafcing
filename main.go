package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/oracle02k/go_raytracing/math3d"
)

func main() {

	nx := 200
	ny := 100

	writeFile, _ := os.OpenFile("test.ppm", os.O_WRONLY|os.O_CREATE, 0600)
	writer := bufio.NewWriter(writeFile)

	writer.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", nx, ny))
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {

			color := math3d.NewVec3(
				float64(i)/float64(nx),
				float64(j)/float64(ny),
				0.2,
			)

			icolor := math3d.Vec3Scale(color, 255.99)
			ir := int32(icolor.R())
			ig := int32(icolor.G())
			ib := int32(icolor.B())
			writer.WriteString(fmt.Sprintf("%d %d %d\n", ir, ig, ib))
		}
	}
	writer.Flush()
	writeFile.Close()
}
