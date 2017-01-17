package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("defer")
	fmt.Println("start")

	os.Exit(255)

	fmt.Println("end")
}
