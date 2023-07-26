package main

import "fmt"

func migratoryBirds(arr []int32) int32 {
	//result := int32(0)
	birds := make(map[int32]int32)

	for i := 0; i < len(arr); i++ {
		birds[arr[i]]++
	}

	maxValue := int32(0)
	maxKey := int32(0)
	for key, value := range birds {
		if value > maxValue {
			maxKey = key
			maxValue = value

		}
	}
	//fmt.Println(checked)
	//fmt.Printf("%d:%d\n", maxKey, maxValue)
	return maxKey
}

func main() {
	fmt.Println(migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4})) //3
	fmt.Println(migratoryBirds([]int32{1, 1, 2, 2, 3}))                   //1
	fmt.Println(migratoryBirds([]int32{1, 4, 4, 4, 5, 3}))                //4
}
