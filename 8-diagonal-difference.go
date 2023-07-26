package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	//Write your code here
	var firstDiagonal int32
	var secondDiagonal int32
	var result int32
	for i := 0; i < len(arr); i++ {
		firstDiagonal += arr[i][i]
		secondDiagonal += arr[i][len(arr)-(i+1)]

	}
	num := math.Abs(float64(firstDiagonal - secondDiagonal))
	result = int32(num)
	return result
}

func main() {
	fmt.Println(diagonalDifference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}})) //15
}
