package main

import "fmt"

func main() {
	var sx, sy, gx, gy int
	fmt.Scanf("%d %d %d %d", &sx, &sy, &gx, &gy)
	r := calc(sx, sy, gx, gy)
	fmt.Printf("%f", r)
}

func calc(sx, sy, gx, gy int) float64 {
	y := sy * gx + gy * sx
	x := gy + sy
	return float64(y) / float64(x)
}
