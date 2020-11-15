package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	nums := make([]int,0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, nextInt())
	}
	maxNum := max(nums)
	d := 0
	ans := 2

	for k := 2; k < 1001; k++ {
		if k > maxNum {
			break
		}
		kd := 0
		for _, n := range nums {
			if n % k == 0 {
				kd++
			}
		}
		if kd == n {
			d = kd
			ans = k
			break
		}
		if kd > d {
			d = kd
			ans = k
		}
	}
	fmt.Printf("%d", ans)
}
