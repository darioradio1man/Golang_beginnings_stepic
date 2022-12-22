package main
import "fmt"

func main() {
    var num int
    _, _ = fmt.Scan(&num)
    fmt.Println(dr(num))

}

func dr(num int) int {
    if num <= 9{
        return num
    } else if num % 9 == 0{
        return 9
    } else{
        return num % 9
    }
}
