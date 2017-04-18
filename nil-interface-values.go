package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	descrive(i)
	i.M()
}

func descrive(i I) {
	fmt.Printf("%v %T", i, i)
}
