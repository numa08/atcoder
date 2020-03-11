package main

import "fmt"

func main() {
	var num, sum int
	fmt.Scanf("%d %d", &num, &sum)
	x, y, z := search(num, sum)
	fmt.Printf("%d %d %d", x, y, z)
}

func search(num, sum int) (x, y, z int) {
	x = -1
	y = -1
	z = -1
	for i := 0; i <= num; i++ {
		for j := 0; j+i <= num; j++ {
			k := num - i - j
			if i+j+k == num && 10000*i+5000*j+1000*k == sum {
				x = i
				y = j
				z = k
				return
			}
		}
	}
	return
}
