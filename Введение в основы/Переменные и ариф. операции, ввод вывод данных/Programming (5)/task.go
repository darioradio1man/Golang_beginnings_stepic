package main
import "fmt"

func main(){
    var degree, minutes, hours int
    fmt.Scan(&degree)
    hours = degree / 30
    minutes = 2 * (degree % 30)
    fmt.Println("It is", hours, "hours", minutes, "minutes.")
}