package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n >= 11 && n <= 14 {
		fmt.Printf("%d korov", n)
	} else {
		temp := n % 10
		if temp == 0 || (temp >= 5 && temp <= 9) {
			fmt.Printf("%d korov", n)
		}
		if temp == 1 {
			fmt.Printf("%d korova", n)
		}
		if temp >= 2 && temp <= 4 {
			fmt.Printf("%d korovy", n)
		}
	}
}
