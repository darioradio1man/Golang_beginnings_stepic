package main

import "fmt"

func main() {
	var num string
	var need string
	var result string

	_, _ = fmt.Scan(&num)
	_, _ = fmt.Scan(&need)

	for _, r := range num {
		if string(r) != need {
			result += string(r)
		}
	}
	fmt.Print(result)
}
