package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")

	panic("panic!!!")
	fmt.Print("Hello ")
}
