package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	//Write your code here
	result := int32(0)
	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}
	return result
}

func main() {
	fmt.Println(simpleArraySum([]int32{1, 2, 3, 4, 10, 11})) //6
}
