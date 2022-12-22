package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	res := int(math.Round(math.Inf(-1)))

	fmt.Scan(&a)
	fmt.Scan(&b)

	for i := a; i <= b; i++ {
		if i%7 == 0 {
			res = i
		}
	}
	if res != int(math.Round(math.Inf(-1))) {
		fmt.Print(res)
	} else {
		fmt.Print("NO")
	}
}
