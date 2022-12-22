workArray := [10]uint8{}
swapArray := [6]uint8{}

for i := 0; i < 10; i++{
    fmt.Scan(&workArray[i])
}
for i := 0; i < 6; i++{
    fmt.Scan(&swapArray[i])
}
for i := 0; i < 5; i+=2{
    workArray[swapArray[i]], workArray[swapArray[i + 1]] = workArray[swapArray[i + 1]], workArray[swapArray[i]]
}
for i := 0; i < 10; i++{
    fmt.Print(workArray[i], " ")
}