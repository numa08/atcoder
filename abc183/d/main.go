package main

import (
	"fmt"
	"math"
)

func main() {
	var n, w int
	fmt.Scanf("%d %d", &n, &w)
	slist := make([]int, n, n)
	tlist := make([]int, n, n)
	plist := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &slist[i], &tlist[i], &plist[i])
	}
	maxT := max(tlist)
	for i := 1; i <= maxT; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if slist[j] <= i && i < tlist[j] {
				sum += plist[j]
			}
		}
		if sum > w {
			fmt.Printf("No")
			return
		}
	}
	fmt.Printf("Yes")
}
func max(nums []int) int {
    if len(nums) == 0 {
        panic("funciton max() requires at least one argument.")
    }
    res := nums[0]
    for i := 0; i < len(nums); i++ {
        res = int(math.Max(float64(res), float64(nums[i])))
    }
    return res
}
