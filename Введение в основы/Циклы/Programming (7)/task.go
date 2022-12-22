package main
import "fmt"

func main(){
    var (
        a, b string
    )
    fmt.Scan(&a, &b)
    for _, x := range a {
        for _, y := range b {
            if (x == y) {
                fmt.Print(string(x) + " ")
            }
        }
    }
}