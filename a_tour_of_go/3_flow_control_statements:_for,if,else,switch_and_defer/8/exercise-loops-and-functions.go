package main

import (
    "fmt"
    "math"
)


func Sqrt(x float64) (float64, int) {
    //z := 1.0
    z := x / 2.0 
    // z := x
    i := 1
    for ; i < 10 ; i++ {
	if math.Abs(x*x - z*z) < 0.0001 {
	    break
	} else {
	    z -= (z*z - x) / (2.0*z)
	}
    }
    return z, i
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}
