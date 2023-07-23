package main

import "fmt"

func gradingStudents(grades []int32) []int32 {
	//writte your code here
	result := []int32{}
	for i := 0; i < len(grades); i++ {
		if grades[i] < 38 {
			result = append(result, grades[i])
			continue
		}
		mod5 := grades[i] % 5
		if mod5 == 3 || mod5 == 4 {
			grades[i] += (5 - mod5)
		} else if mod5 == 2 || mod5 == 1 {
			grades[i] = grades[i]
		}
		result = append(result, grades[i])
	}
	return result
}

func main() {
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33})) //75 67 40 33
}
