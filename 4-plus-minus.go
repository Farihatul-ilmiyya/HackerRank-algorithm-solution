package main

import "fmt"

func plusMinus(arr []int32) {
	//write your code here
	var negative float64
	var positive float64
	var zero float64
	n := float64(len(arr))
	//
	for i := 0; i < len(arr); i++ {
		if arr[i] < 0 {
			negative = negative + 1
		} else if arr[i] > 0 {
			positive += 1
		} else {
			zero += 1
		}

	}
	fmt.Printf("%.6f\n", positive/n)
	fmt.Printf("%.6f\n", negative/n)
	fmt.Printf("%.6f\n", zero/n)

}

func main() {
	plusMinus([]int32{-1, 3, -9, 0, 4, 1}) //0.500000 0.333333 0.166667
}
