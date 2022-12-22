package main
import "fmt"

func main() {
    var x, y float64
    fmt.Scan(&x)
    y = x * x
    if x <= 0 {
        fmt.Printf("число %.2f не подходит", x)
    } else if x <= 10000 {
        fmt.Printf("%.4f", y)
    } else {
        fmt.Printf("%e", x)
    }
}