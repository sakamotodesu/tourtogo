package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Vertex{1, 2})
	v.X = 8
	fmt.Println(v.X)
}
