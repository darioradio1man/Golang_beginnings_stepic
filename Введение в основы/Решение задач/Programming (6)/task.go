package main
import "fmt"

func main() {
    var n int
    res:=0
    fmt.Scanln(&n)
    array := make([]int, n)
    for i:=0; i < n; i++{
        fmt.Scan(&array[i])
    }
    for i:=0; i < n; i++{
        if array[i] == 0{
            res++
        }
    }
    fmt.Print(res)
}