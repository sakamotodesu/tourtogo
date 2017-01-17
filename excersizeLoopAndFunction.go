package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z = float64(3)
	for n := 1; n < 10; n++ {
		z = zzz(z, x)
		fmt.Println(z)
	}
	return z
}

func zzz(z, x float64) float64 {
	return z - ((z*z - x) / 2 * z)
}

func main() {
	fmt.Println(Sqrt(2))
}
