package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 2, 1, 2, 1, 5, 2, 2, 1, 5, 1, 2, 1, 14, 1, 5, 1, 5, 2, 2, 1, 15, 2, 2, 5, 4, 1, 4, 1, 51}
	var k int
	fmt.Scanf("%d", &k)
	fmt.Printf("%d", arr[k-1])
}
