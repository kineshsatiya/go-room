package main

import (
    "fmt"
    "math"
)

func Sqrt(x float64) float64 {
    var z float64 = 1
    var sqrt float64 = 0
    for {
        z = z - (z * z - x)/(2*z)
        if math.Abs(sqrt - z) < 1e-15 {
            break
        }
        sqrt = z
    }
    return sqrt
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(math.Sqrt(2))
}
