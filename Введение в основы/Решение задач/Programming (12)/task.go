package main

import "fmt"

func main() {
	var num, i, nNext, nFirst int

	fmt.Scan(&num)
	for nSecond := 1; num >= nNext; i++ {
		nNext = nFirst + nSecond
		nSecond, nFirst = nNext, nSecond
	}
	switch {
	case nFirst == num:
		fmt.Println(i)
	default:
		fmt.Print(-1)
	}
}
