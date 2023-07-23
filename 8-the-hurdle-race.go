package main

import "fmt"

func hurdleRace(k int32, height []int32) int32 {
	//write your code here
	var result int32
	max := height[0]
	for i := 0; i < len(height); i++ {
		if height[i] > max {
			max = height[i]
		}
	}
	if max < k {
		result = 0
	} else {
		result = max - k
	}
	return result
}

func main() {
	fmt.Println(hurdleRace(1, []int32{1, 2, 3, 3, 2})) // 2
	fmt.Println(hurdleRace(4, []int32{1, 6, 3, 5, 2})) // 2
	fmt.Println(hurdleRace(7, []int32{2, 5, 4, 5, 2})) // 0

}
