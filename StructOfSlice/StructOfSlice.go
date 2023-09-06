//Написать Struct of slice without using slice ✅
// Struct should include cap, len, pointer to arr ✅
// cap of array should be assigned by const ✅

// major methods should be add, delete, and binSearch [CRUD]

package main

import (
	"fmt"
)

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

// Бабл сорт implemented ✅
// time complexity n^2 - 2 цикла, space O(1) - не требует лишней памяти, все объявления по 1
func (s *StructOfArr) sort() {
	for i := 0; i < s.length-1; i++ {
		for j := 0; j < s.length-i-1; j++ {
			if s.arr[j] > s.arr[j+1] {
				s.arr[j], s.arr[j+1] = s.arr[j+1], s.arr[j]
			}
		}
	}
}

//

// Добавить Create = add ✅
// что бы добавить новую цифру нужно чекнуть, заполнен ли массив return true false ✅
// Нужен метод исползующий стракт ✅
func (s *StructOfArr) add(num int) bool {
	if s.length >= s.cap {
		return false
	}

	s.arr[s.length] = num
	s.length++
	return true
}

func main() {

	arr := [arrCap]int{4, 5, 6, 7}
	arry := &StructOfArr{
		length: 3,
		cap:    arrCap,
		arr:    &arr,
	}

	var cifra int
	fmt.Print("Введите целую цифру от 0-9 для поиска в массиве\n")
	fmt.Scan(&cifra)
	//

	//
	res := binSearch(arry.arr, arry.length, cifra)

	if res != -1 {
		fmt.Printf("\nВаша цифра %d находится в индексе %d\n", cifra, res)
	} else {
		fmt.Printf("Ваша цифра %d не найдена\n", cifra)
	}

	fmt.Println(arr)
	fmt.Printf("\nЧто бы добавить цифру к массиву напишите цифру 1,\nчто бы удалить, цифру 0\n")
	//
	var addOrDelete int
	fmt.Scan(&addOrDelete)
	//

	if addOrDelete == 1 {

		fmt.Printf("Напишите цифру которую хотите добавить\n")
		var newInt int
		fmt.Scan(&newInt)
		//
		if arry.add(newInt) {
			fmt.Println("Массив полон, нельзя добавить цифру")
		} else {
			fmt.Printf("Цифра %d добавлена в массив\n", newInt)
			fmt.Println(arry.arr)
		}
		//
		addNumArr := arry.add(newInt)
		// fmt.Println(addNumArr)
		if addNumArr {
			fmt.Printf("Цифра %d добавлена в массив\n", newInt)
			fmt.Println(arry.arr)
		}
	}
	// Добавить Create = add 🔍
	// Delete = delete 🔍
	// Update 🔍
	// Read 🔍

}
