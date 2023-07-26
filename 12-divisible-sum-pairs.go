package main

import "fmt"

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var result int32
	var num2 = []int32{}
	for i := 0; i < len(ar); i++ {
		for j := 0; j < len(ar); j++ {
			if i == j {
				break
			}
			num1 := int32(ar[i] + ar[j])
			if num1%k == 0 {
				num2 = append(num2, num1)
			}

		}

	}
	fmt.Println(num2)
	result = int32(len(num2))
	return result
}

func main() {
	fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))

}
