package main

import "fmt"

type User struct {
    Name string
    Age int
}

func main() {
    var userA User
    userA.Name = "太郎"
    userA.Age = 30
    fmt.Println(userA)
}
