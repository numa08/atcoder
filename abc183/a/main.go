package main

import "fmt"

func main() {
	var i int
	fmt.Scanf("%d", &i)
	r := reLU(i)
	fmt.Printf("%d", r)
}

func reLU(x int) int {
	if x < 0 {
		return 0
	}
	return x
}
