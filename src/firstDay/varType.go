package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func emut() {
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func main() {
	emut()
}
