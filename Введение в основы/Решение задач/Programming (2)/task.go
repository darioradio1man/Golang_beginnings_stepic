package main
import "fmt"

func main(){
    var (
        k, hours, minutes int
    )
    
    fmt.Scan(&k)
    hours = k / 3600
    k %= 3600
    minutes = k / 60
    fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}