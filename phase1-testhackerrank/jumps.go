package main

import "fmt"

func jumps(flagHeight int32, bigJump int32) int32 {
	var result int32
	jumping := int32(0)
	//minimumJump := int32(1)

	for jumping < flagHeight {
		// Coba lompatan besar (bigJump) jika masih memungkinkan.
		if jumping+bigJump <= flagHeight {
			jumping += bigJump
			result++
		} else {
			// Jika tidak, lakukan lompatan kecil (ketinggian 1).
			jumping++
			result++
		}
	}

	return result
}

func main() {
	fmt.Println(jumps(3, 1)) //3
	fmt.Println(jumps(3, 2)) //2
	fmt.Println(jumps(3, 3)) //1
	fmt.Println(jumps(8, 3)) //4
}
