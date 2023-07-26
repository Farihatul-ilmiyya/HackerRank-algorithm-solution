package main

import (
	"fmt"
	"math"
)

func catAndMouse(x int32, y int32, z int32) string {
	position1 := math.Abs(float64(z - x))
	position2 := math.Abs(float64(z - y))
	var result string
	if position1 > position2 {
		result += fmt.Sprintf("Cat B")
	} else if position1 < position2 {
		result += fmt.Sprintf("Cat A")
	} else {
		result += fmt.Sprintf("Mouse C")
	}

	return result
}

func main() {
	fmt.Println(catAndMouse(1, 2, 3)) // Cat B
	fmt.Println(catAndMouse(1, 3, 2)) // Mouse C
}
