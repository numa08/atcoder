package main

import "fmt"

func main() {
	var follower, current int
	fmt.Scanf("%d %d", &follower, &current)
	max := 2 * follower + 100
	latest := max - current
	fmt.Printf("%d", latest)
}
