package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	//write your code here
	var result = []int32{}
	max := scores[0]
	var highest = []int32{}
	min := scores[0]
	var lowest = []int32{}
	for i := 0; i < len(scores); i++ {
		//fmt.Println(scores[i])
		if scores[i] > max {
			max = scores[i]
			highest = append(highest, max)

		}
		if scores[i] < min {
			min = scores[i]
			lowest = append(lowest, min)
		}

	}
	//fmt.Println(highest)
	//fmt.Println(lowest)
	num1 := len(highest)
	num2 := len(lowest)
	result = append(result, int32(num1), int32(num2))
	return result
}

func main() {
	fmt.Println(breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42})) //[4 0]
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))       //[2 4]

}
