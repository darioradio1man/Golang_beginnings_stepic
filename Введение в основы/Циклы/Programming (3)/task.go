package main

import "fmt"

func main()  {
	var n int
	var flag int = 1
	_, _ = fmt.Scan(&n)
	maximum := n
	// считываем числа пока не будет введен 0
	for fmt.Scan(&n); n != 0; fmt.Scan(&n){
		if n > maximum{
			maximum = n
			flag = 0
		}
		if n == maximum{
			flag++
		}
	}
	fmt.Print(flag)
}