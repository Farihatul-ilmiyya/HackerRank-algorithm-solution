package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	max := int32(0)
	var result int32
	var num = []int32{}
	for i := 0; i < len(candles); i++ {
		if candles[i] > max {
			max = candles[i]
			num = []int32{max}
		} else if candles[i] == max {
			num = append(num, max)
		}
	}

	result = int32(len(num))
	return result
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3})) //2
}
