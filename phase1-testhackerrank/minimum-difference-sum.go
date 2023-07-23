package main

import (
	"fmt"
	"math"
	"sort"
)

func minDiff(arr []int32) int32 {
	//write your code here
	var result int32
	num := arr
	sort.Slice(num, func(i, j int) bool {
		return num[i] < num[j]
	})
	// for _, v := range arr {
	// 	fmt.Println(v)
	// }
	//fmt.Println(num)
	counter := int32(0)
	for i := 0; i < len(num)-1; i++ {
		counter = int32(math.Abs(float64(num[i] - num[i+1])))
		result += counter
		fmt.Println(counter)
	}
	return result
}

func main() {
	fmt.Println(minDiff([]int32{1, 3, 3, 2, 4})) //3
	fmt.Println(minDiff([]int32{5, 1, 3, 7, 3})) //6

}
