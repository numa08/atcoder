package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	p := int(float64(b) / 0.1)
	ans := -1
	for i := 0; i < 10; i++ {
		if int(float64(p+i)*0.08) == a {
			ans = p + i
			break
		}
	}
	fmt.Printf("%d", ans)
}
