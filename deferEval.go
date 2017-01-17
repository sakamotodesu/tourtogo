package main

import (
	"fmt"
)

func main() {
	n := 1
	defer fmt.Printf("n : %d", n)
	n *= 100
	fmt.Printf("n:100 %d", n)
}
