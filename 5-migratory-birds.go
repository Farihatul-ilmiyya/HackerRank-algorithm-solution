package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	//result := int32(0)
	checked := make(map[int32]int32)

	for i := 0; i < len(arr); i++ {
		value, exist := checked[arr[i]]
		fmt.Println(value, exist)
		if exist {
			checked[arr[i]]++
		} else {
			checked[arr[i]] = 1
		}
	}
	var maxValue int32
	var minKey int32
	for key := range checked {
		minKey = key
		break
	}

	for key, value := range checked {
		if value > maxValue {
			maxValue = value
		}
		if value == maxValue && key < minKey {
			minKey = key
		}

	}
	//fmt.Println(checked)
	return minKey
}

func main() {
	fmt.Println(migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4})) //3
	//fmt.Println(migratoryBirds([]int32{1, 1, 2, 2, 3}))                   //1
	//fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))
}
