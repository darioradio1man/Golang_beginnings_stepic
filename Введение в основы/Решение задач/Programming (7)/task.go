package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	res := 0
	var min int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	sort.Ints(arr)
	min = arr[0]
	for i := 0; i < n; i++ {
		if min == arr[i] {
			res += 1
		}
	}
	fmt.Println(res)
}
