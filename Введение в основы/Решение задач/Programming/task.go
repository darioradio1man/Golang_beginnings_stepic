package main 
import "fmt"

func main(){
    var x, x1, x2, x3 int
    fmt.Scan(&x)
    x1 = x / 100
    x2 = (x % 100) / 10
    x3 = x % 10
    fmt.Println(x1 + x2 + x3)
}