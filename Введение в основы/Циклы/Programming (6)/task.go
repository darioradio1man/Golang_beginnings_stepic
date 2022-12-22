package main
import "fmt"

func main(){
    var x, p, y, d int
    fmt.Scan(&x, &p, &y)
    for ; x < y; d++{
        x += int((x * p) / 100.0)
    }
    fmt.Println(d)
}