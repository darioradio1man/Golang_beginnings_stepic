package main
import "fmt"

func main() {
    var n int
    fmt.Scanln(&n)
    array := make([]int, n)
    for i:=0; i < n; i++{
        fmt.Scan(&array[i])
    }
    result := 0
    for i:=0; i < n; i++{
        if array[i] > 0 {
            result++
        }
    }
    fmt.Println(result)
}