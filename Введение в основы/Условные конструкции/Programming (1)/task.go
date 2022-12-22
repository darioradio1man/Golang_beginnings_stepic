package main

import "fmt"

func main()  {
	var num int

	_,_ = fmt.Scan(&num)
	x1 := num / 100
	x2 := (num % 100) / 10
	x3 := num % 10
	if x1 != x2 && x2 != x3 && x3 != x1{
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}