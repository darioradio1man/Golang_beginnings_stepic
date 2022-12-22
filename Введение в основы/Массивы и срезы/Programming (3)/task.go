package main
import "fmt"

func main() {
    var n int
    fmt.Scanln(&n)
    array := make([]int, n)
    for i:=0; i < n; i++{
        fmt.Scan(&array[i])
    }
    for i:=0; i < n; i+=2{
        fmt.Print(array[i], " ")
    }
}