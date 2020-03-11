package main

import (
	"fmt"
)

func main()  {
	var c int
	fmt.Scanf("%d", &c)
	var l = make([]int, 200 ,200)
	for i := 0; i < c; i++ {
		fmt.Scanf("%d", &l[i])
	}
	check := func (l []int) int {
		d := 0
		for {
			for i, k := range l {
				if k % 2 != 0 {
					return d
				}
				l[i] = k / 2
			}
			d++
		}
	}
	ans := check(l)
	fmt.Printf("%d", ans)
}