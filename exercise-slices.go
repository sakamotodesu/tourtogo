package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var s = make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s[i] = make([]uint8, dx)
	}

	for x := 0; x < dx; x++ {
		for y := 0; y < dy; y++ {
			s[x][y] = uint8(omoshiroi(x, y))
		}
	}
	return s
}

func omoshiroi(x, y int) int {
	return x * y
}

func main() {
	pic.Show(Pic)
}
