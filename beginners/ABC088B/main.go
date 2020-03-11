package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	cards := make(sort.IntSlice, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &cards[i])
	}

	alice := 0
	bob := 0
	sort.Sort(sort.Reverse(cards))

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice += cards[i]
		} else {
			bob += cards[i]
		}
	}
	ans := alice - bob
	fmt.Printf("%d", ans)
}
