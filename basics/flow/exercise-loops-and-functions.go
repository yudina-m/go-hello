package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z_prev := 0.0
	delta := 0.001
	for delta >= 0.001 {
		z_prev = z
		z -= (z*z - x) / 2 * z

		delta = z_prev - z

		if delta < 0 {
			delta *= -1
		}

		if delta > 0.001 {
			fmt.Println(z, delta)
		}
	}
	return z
}

func main()  {
	fmt.Println("===>", Sqrt(2), math.Sqrt(2))
}
