package main
import "fmt"

const Pi = 3.14

const (
    URL = "http://xxx.co.jp"
    SiteName = "test"
)

const (
    A = 1
    B
    C
    D = "A"
    E
    F
)

func main() {
    fmt.Println(Pi)

    fmt.Println(URL,SiteName)

    fmt.Println(A,B,C,D,E,F)
}
