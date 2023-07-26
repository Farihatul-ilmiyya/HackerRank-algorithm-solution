package main

import "fmt"

func bonAppetit(bill []int32, k int32, b int32) {
	bActual := int32(0)
	bAnna := int32(0)
	bCharged := b
	for i := 0; i < len(bill); i++ {
		// for j := 0; j<len(bill); j++
		//fmt.Println(i)
		if bill[i] != bill[k] {
			bAnna += bill[i]
		}
		bActual += bill[i]

	}
	//yang harus anna bayar
	totalAnna := bAnna / 2
	//yang harus anna bayar jika dia makan semuanya
	totalAll := bActual / 2
	//kembalian yang anna terima
	charged := bCharged - totalAnna
	//jika anna makan semuanya maka
	if k == 0 {
		fmt.Println(bCharged - totalAll)
	} else {
		if charged != 0 {
			fmt.Println(charged)
		} else if charged == 0 {
			fmt.Println("Bon Appetit")
		}
	}
	//fmt.Println(bill[k])
}

func main() {
	bonAppetit([]int32{3, 10, 2, 9}, 1, 12) //5
	bonAppetit([]int32{3, 10, 2, 9}, 1, 7)  //Bon Appetit
	bonAppetit([]int32{3, 10, 2, 9}, 0, 7)  //Bon Appetit
}
