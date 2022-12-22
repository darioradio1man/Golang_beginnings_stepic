package main
import "fmt"

func main() {
    var a, b int
    sum := 0
    fmt.Scan(&a, &b)
    for i := a; i <= b; i++{
        sum += i
    }
    fmt.Println(sum)
}