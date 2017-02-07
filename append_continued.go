package main

import "fmt"

func main() {
	array := [6]int{2, 3, 5, 7, 11, 13}
	slice := array[1:4]

	fmt.Println(array)
	fmt.Println(slice)

	slice = append(slice, 999)
	fmt.Println(array)
	fmt.Println(slice)

}
