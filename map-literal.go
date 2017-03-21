package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		123, 234,
	},
	"Google": Vertex{
		234, 345,
	},
}

func main() {
	fmt.Println(m)
}
