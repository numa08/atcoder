package main

import (
	"fmt"
)

func main()  {
	var s string
	fmt.Scanf("%s", &s)
	sum := 0
	if s[0] == '1' {
		sum++
	}
	if s[1] == '1' {
		sum++
	}
	if s[2] == '1' {
		sum++
	}
	fmt.Printf("%d", sum)
}
