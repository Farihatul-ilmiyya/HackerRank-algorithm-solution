package main

import "fmt"

func utopianTree(n int32) int32 {
	for i := int32(0); i <= n; i++ {
		if n == 0 {
			return 1
		} else if n%2 == 1 {
			return 2 * utopianTree(n-1)
		} else if n%2 == 0 {
			return 1 + utopianTree(n-1)
		}
	return 0
}

func main() {
	fmt.Println(utopianTree(0)) //1
	fmt.Println(utopianTree(1)) //2
	fmt.Println(utopianTree(3)) //6
	fmt.Println(utopianTree(4)) //7
	fmt.Println(utopianTree(5)) //14

}
