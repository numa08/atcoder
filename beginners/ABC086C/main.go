package main

import (
	"fmt"
	"math"
)

func main() {
	cx := 0
	cy := 0
	ct := 0

	var n, t, x, y int
	fmt.Scanf("%d", &n)
	can := true
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &t, &x, &y)
		dx := int(math.Abs(float64(x) - float64(cx)))
		dy := int(math.Abs(float64(y) - float64(cy)))
		dt := t - ct
		d := dx + dy

		can = d <= dt && d%2 == dt%2
		if !can {
			break
		}
		cx = x
		cy = y
		ct = t
	}

	if can {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
