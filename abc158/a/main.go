package main

import "fmt"

func main() {
	length := 3
	var stations string
	fmt.Scanf("%s", &stations)
	prevStation := ""
	needs := false
	for i := 0; i < length; i++ {
		current := string(stations[i])
		if prevStation == "" {
			prevStation = current
			continue
		}
		needs = prevStation != current
		if needs {
			break
		}
		prevStation = current
	}
	if needs {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}
