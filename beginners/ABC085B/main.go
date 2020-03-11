package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	mochis := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &mochis[i])
	}

	mochiMap := make(map[int]bool)
	uniqueMochi := []int{}

	for _, mochi := range mochis {
		if !mochiMap[mochi] {
			uniqueMochi = append(uniqueMochi, mochi)
			mochiMap[mochi] = true
		}
	}

	ans := len(uniqueMochi)
	fmt.Printf("%d", ans)
}
