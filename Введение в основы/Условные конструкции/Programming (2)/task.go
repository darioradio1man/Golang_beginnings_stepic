package main
import "fmt"

func main(){
    var x int
    fmt.Scan(&x)
    for x >= 10 {
        x = x / 10
    }
    fmt.Println(x)
} 