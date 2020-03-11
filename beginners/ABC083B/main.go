package main

import (
	"fmt"
)

func main() {
	var n, a, b int
	ans := 0
	fmt.Scanf("%d %d %d", &n, &a, &b)
	for i := 1; i <= n; i++ {
		s := sum(i)
		if s >= a && s <= b {
			ans += i
		}
	}
	fmt.Printf("%d", ans)
}

func sum(target int) int {
	sum := 0
	t := target
	for {
		if t < 10 {
			return sum + t
		}
		sum += t % 10
		t /= 10
	}
}
