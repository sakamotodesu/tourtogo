package main

import "fmt"

func main() {
	var i interface{}
	descrive(i)

	i = 42
	descrive(i)

	i = "hello"
	descrive(i)
}

func descrive(i interface{}) {
	fmt.Printf("%v %T", i, i)
}
