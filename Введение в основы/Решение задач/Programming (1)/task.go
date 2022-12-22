package main 
import (
    "fmt"
    "strconv"
)

func main(){
    var (
        x int
        x1, x2, x3 string
    )
    
    fmt.Scan(&x)
    x1 = strconv.Itoa(x / 100)
    x2 = strconv.Itoa((x % 100) / 10)
    x3 = strconv.Itoa(x % 10)
    fmt.Println(x3 + x2 + x1)
}