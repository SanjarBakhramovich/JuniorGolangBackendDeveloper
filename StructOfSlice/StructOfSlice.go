//Написать Struct of slice without using slice ✅
// Struct should include cap, len, pointer to arr ✅
// cap of array should be assigned by const ✅

// major methods should be add, delete, and binSearch [CRUD]

package main

import (
	"JuniorSpecMod/StructOfSlice/slicePckg"
	"fmt"
)

// MAIN
func main() {

	arr := [slicePckg.ArrCap]int{4, 5, 6, 7}
	arry := &slicePckg.StructOfArr{
		Length: 3,
		Cap:    slicePckg.ArrCap,
		Arr:    &arr,
	}

	var cifra int
	fmt.Print("Введите целую цифру от 0-9 для поиска в массиве\n")
	fmt.Scan(&cifra)
	//

	//
	res := slicePckg.BinSearch(arry.Arr, arry.Length, cifra)

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
		if arry.Add(newInt) {
			fmt.Println("Массив полон, нельзя добавить цифру")
		} else {
			fmt.Printf("Цифра %d добавлена в массив\n", newInt)
			fmt.Println(arry.Arr)
		}
		//
		addNumArr := arry.Add(newInt)
		// fmt.Println(addNumArr)
		if addNumArr {
			fmt.Printf("Цифра %d добавлена в массив\n", newInt)
			fmt.Println(arry.Arr)
		}
	}
	// Добавить Create = add 🔍
	// Delete = delete 🔍
	// Update 🔍
	// Read 🔍

}

// Мэйн зделать отдельным файлом
// Конструктор (вызывать конструктор через мэйн)
// package slice отдельный создать
// импортировать slice в main
// пользователь пишет getSlice
// Методы зделать публичными
