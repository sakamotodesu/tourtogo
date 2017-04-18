package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v ", v, v*2)
	case string:
		fmt.Printf("%v is %v byte long", v, len(v))
	default:
		fmt.Printf("I don't know about type %T !", v)
	}
}
func main() {
	do(21)
	do("hello")
	do(true)

}
