package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x >= 0 {
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
        return z, nil
	} else {
	    return x, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
