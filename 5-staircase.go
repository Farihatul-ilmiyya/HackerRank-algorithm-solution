package main

import "fmt"

func staircase(n int32) {
	// write your code here
	for rows := 1; rows <= int(n); rows++ {
		for space := 1; space <= int(n)-rows; space++ {
			fmt.Print(" ")
		}
		for col := 1; col <= rows; col++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func main() {
	staircase(6)
}
