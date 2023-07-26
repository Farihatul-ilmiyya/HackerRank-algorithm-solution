package main

import "fmt"

func designerPdfViewer(h []int32, word string) int32 {
	//write your code here
	wide := int32(1)
	var result int32
	atoz := make(map[rune]int32)
	for i := 0; i < len(h); i++ {
		//a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z := h[0sampai25]
		atoz[rune('a'+i)] = h[i]
	}
	maxValue := int32(0)
	for _, char := range word {
		if value, exist := atoz[char]; exist {
			//fmt.Printf("%c: %d\n", char, value)
			if value > maxValue {
				maxValue = value
			}
		}
	}

	//fmt.Println(maxValue)
	result = maxValue * int32(len(word)) * wide
	return result
}

func main() {
	fmt.Println(designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 1, 1, 5, 5, 1, 5, 2, 5, 5, 5, 5, 5, 5}, "torn")) //8
	fmt.Println(designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 1, 1, 5, 5, 1, 5, 2, 5, 5, 5, 5, 5, 5}, "abc"))  //9
	fmt.Println(designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 1, 1, 5, 5, 1, 5, 2, 5, 5, 5, 5, 5, 5}, "zaba")) //28
}
