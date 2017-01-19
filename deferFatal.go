package main

import (
	"fmt"
	"log"
)

func main() {
	defer fmt.Println("defe")

	fmt.Println("start")

	log.Fatal("fatal!")

	fmt.Println("end")
}
