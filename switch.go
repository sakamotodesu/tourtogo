package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on")

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
	default:
		fmt.Printf("%s.", os)
	}
}
