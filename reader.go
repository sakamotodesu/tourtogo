package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello reader!")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("%v %v %v", n, err, b)
		fmt.Println("")
		fmt.Printf("%q", b[:n])
		if err == io.EOF {
			break
		}
	}

}
