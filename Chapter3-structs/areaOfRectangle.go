package main

import "fmt"

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	var rect rectangle
	rect.width = 200
	rect.height = 100

	fmt.Printf("Area is equal to: %f \n", rect.Area())

}
