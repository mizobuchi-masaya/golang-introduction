package main

import "fmt"

func Double(i int) {
    i = i * 2
}

func Doublev2(i *int) {
    *i = *i * 2
}

func main() {
    var n int = 10000
    fmt. Println(n)
    fmt. Println(&n)

    Double(n)
    fmt.Println(n)

    var p *int = &n
    fmt.Println(p)
    fmt.Println(*p)

    Doublev2(&n)
    fmt.Println(n)
}
