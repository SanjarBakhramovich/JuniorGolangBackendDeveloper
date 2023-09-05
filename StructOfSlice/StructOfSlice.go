//Написать Struct of slice without using slice
// Struct should include cap, len, pointer to arr

// cap of array should be assigned by const

// major methods should be add, delete, and binSearch [CRUD]

package main

import "fmt"

const arrCap = 4

type StructOfArr struct {
	length int
	cap    int
	arr    *[arrCap]int
}

func binSearch(arr *[arrCap]int, length, target int) int {
	low := 0
	high := length - 1

	for low <= high {
		mid := low + (high-low)/2 //идёт в центр

		if arr[mid] == target {
			return mid // если нашёл возвращает
		} else if arr[mid] < target {
			low = mid + 1 // идёт вправо
		} else {
			high = mid - 1 // идёт влево
		}
	}
	// выход из лупа
	return -1 //если цифра не найдена возвращает -1 (ошибку)
}

func main() {

	arr := [arrCap]int{4, 5, 6, 7}
	arry := &StructOfArr{
		length: 3,
		cap:    arrCap,
		arr:    &arr,
	}

	var cifra int
	fmt.Print("Введите целую цифру от 0-9\n")
	fmt.Scan(&cifra)

	res := binSearch(arry.arr, arry.length, cifra)

	if res != -1 {
		fmt.Printf("\nВаша цифра %d находится в индексе %d\n", cifra, res)
	} else {
		fmt.Printf("Ваша цифра %d не найдена\n", cifra)
	}
}
