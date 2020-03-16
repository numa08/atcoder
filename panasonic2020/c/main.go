package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scanf("%v %v %v", &a, &b, &c)

	ans := a+b-c < -2*math.Sqrt(a*b)
	if ans {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
