package main

import (
	"fmt"
)

func closedPaths(number int32) int32 {
	totalPaths := int32(0)
	for number != 0 {
		digit := number % 10
		if digit == 0 || digit == 4 || digit == 6 || digit == 9 {
			totalPaths += 1
		} else if digit == 8 {
			totalPaths += 2
		}
		number /= 10
	}
	return totalPaths
}

func main() {
	fmt.Println(closedPaths(649578)) //
}
