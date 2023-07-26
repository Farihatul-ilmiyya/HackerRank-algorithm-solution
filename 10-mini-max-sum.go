package main

import "fmt"

func miniMaxSum(arr []int32) {
	//write your code here
	var num = []int64{}
	var min int64
	var max int64

	for i := 0; i < len(arr); i++ {
		//fmt.Println(arr[i])
		counter := int64(0)
		for j := 0; j < len(arr); j++ {
			if i == j {
				continue
			}
			counter += int64(arr[j])
		}
		num = append(num, counter)
	}
	max = num[0]
	min = num[0]
	for _, value := range num {
		if value > max {
			max = value
		}
		if value < min {
			min = value
		}
	}

	fmt.Println(min, max)
}

func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})
	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})
}
