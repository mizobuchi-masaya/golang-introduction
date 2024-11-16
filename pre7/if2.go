package main

import "fmt"

func main() {
    a := 10
    if a == 20 {
	fmt.Println("twenty")
    } else if a == 10 {
	fmt.Println("ten")
    } else {
	fmt.Println("Nothing")
    }
}
