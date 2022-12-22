package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	var result string

	_, _ = fmt.Scan(&num)
	for i := 0; i < num; i++ {
		x := int(math.Exp2(float64(i)))
		if x > num {
			break
		}
		result += fmt.Sprintf("%d ", x)
	}
	fmt.Print(result)
}
