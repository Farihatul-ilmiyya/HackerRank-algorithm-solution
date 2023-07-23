package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	//write your code
	var result = []int32{}
	counterA := 0
	counterB := 0
	for i := 0; i < len(a); i++ {
		//fmt.Println(a[i])
		//fmt.Println(b[i])
		if a[i] > b[i] {
			counterA++
		} else if a[i] < b[i] {
			counterB++
		} else if a[i] == b[0] {
			counterA = counterA + 0
			counterB = counterB + 0
		}
		//fmt.Println("->", counterA)
		//fmt.Println("*", counterB)

	}
	result = append(result, int32(counterA), int32(counterB))
	return result
}

func main() {
	fmt.Println(compareTriplets([]int32{17, 28, 30}, []int32{99, 16, 8})) // [2, 1]
}
