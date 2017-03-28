package main

import "fmt"

func fibo() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		next := f1 + f2
		f1 = f2
		f2 = next
		return next
	}
}

func main() {
	f := fibo()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
