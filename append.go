package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	s = append(s, 2, 3, 4)
	printSlice(s)

}

func printSlice(x []int) {
	fmt.Printf("%d %d %v %p", len(x), cap(x), x, x)
}
