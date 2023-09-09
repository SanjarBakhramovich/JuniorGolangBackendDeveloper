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
		fmt.Printf("Ваша цифра %d не найдена в массиве\n", cifra)
	}

	var branchLogic string
	fmt.Printf("Выберите и напишите одну из операции \n[getSlice], [addNum], [delNum] или [update]: ")
	fmt.Scan(&branchLogic)
	//

	// Создать 4 метода для этих операции
	switch branchLogic {
	case "getSlice":
		fmt.Println(arry.GetSlice())

	case "addNum":
		var numToAdd int
		fmt.Print("Введите число для добавления: ")
		fmt.Scan(&numToAdd)
		if arry.Add(numToAdd) {
			fmt.Println("Число успешно добавлено")
		} else {
			fmt.Println("Не удалось добавить число. Массив заполнен.")
		}
	case "delNum":
		var numDel int
		fmt.Print("Введите число для удаления: ")
		fmt.Scan(&numDel)
		if arry.DeleteVal(numDel) {
			fmt.Println("Число успешно удалено")
		} else {
			fmt.Println("Не удалось удалить число. Число не найдено.")
		}

	case "update":
		var idxUpdate, newVal int
		fmt.Print("Введите индекс для обновления: ")
		fmt.Scan(&idxUpdate)
		fmt.Print("Введите новое значение: ")
		fmt.Scan(&newVal)
		if arry.Update(idxUpdate, newVal) {
			fmt.Println("Число успешно обновлено")
		} else {
			fmt.Println("Не удалось обновить число. Неверный индекс.")
		}
	default:
		fmt.Println("Вы не выбрали один из вариантов")
	}

	// Мэйн зделать отдельным файлом✅
	// Конструктор (вызывать конструктор через мэйн)
	// package slice отдельный создать✅
	// импортировать slice в main✅
	// пользователь пишет getSlice
	// Методы зделать публичными✅
}

// Нужно пофиксить метод GetSlice, ты перевыделяешь память, чтобы обратиться к существующей
// В функции Update нужно написать логику по увелечению массива в 2 раза
// при этом поменяв массив на интерфейс в самой структуре
