package main
import "fmt"

func main(){
    var a, a1, a2, a3, a4, a5, a6 int
    fmt.Scan(&a)
    a1 = a / 100000
    a2 = (a % 100000) / 10000
    a3 = (a % 10000) / 1000
    a4 = (a % 1000) / 100
    a5 = (a % 100) / 10
    a6 = a % 10
    if a1 + a2 + a3 == a4 + a5 + a6 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}