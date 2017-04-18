package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	var f = fmt.Sprint(float64(e))
	return fmt.Sprintf("cannot sqrt negative number : %v", f)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	var z = float64(1)
	for n := 1; n < 10; n++ {
		z = zzz(z, x)
		fmt.Println(z)
	}
	return z, nil
}

func zzz(z, x float64) float64 {
	return z - ((z*z - x) / 2 * z)
}

func main() {

	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))

}
