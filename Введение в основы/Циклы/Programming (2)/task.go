package main
import "fmt"

func main(){
    var n, ent int
    sum := 0
    fmt.Scan(&n)
    for i := 1; i <= n; i++{
        fmt.Scan(&ent)
        if ent % 8 == 0 && ent < 100 && ent > 9 { 
            sum = sum + ent 
        } 
    } 
    fmt.Print(sum) 
}