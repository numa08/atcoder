package main

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

func main() {
	var re = regexp.MustCompile(`(?s)^(dream(er)?|eraser?)*$`)
	var str string
	fmt.Scanf("%s", &str)

	t := re.ReplaceAllString(str, "")

	if utf8.RuneCountInString(t) == 0 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
