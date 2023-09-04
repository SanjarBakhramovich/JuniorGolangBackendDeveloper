//Написать Struct of slice without using slice
// Struct should include cap, len, pointer to arr

// cap of array should be assigned by const

// major methods should be add, delete, and binSearch [CRUD]

package main

import "fmt"

func main() {

	const arrCap = 4

	type StructOfArr struct {
		length int
		cap    int
		arr    *[arrCap]int
	}

	arr := [arrCap]int{4, 5, 6, 7}
	arry := &StructOfArr{
		length: 3,
		cap:    arrCap,
		arr:    &arr,
	}

	fmt.Println(arry) //&{3 4 [0xc00012c780]}
}
