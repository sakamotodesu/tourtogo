package main

import "fmt"
import "math"

func sq(x float64) string {
	if x < 0 {
		return sq(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main() {
	fmt.Println(sq(2), sq(-4))
}
