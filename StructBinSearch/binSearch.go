package main

import (
	"fmt"
)

const MaxCapacity = 10

// Структура DynArr представляет динамический массив, основанный на связанных списках.
type DynArr struct {
	head *Node
}

// Структура Node представляет узел связанного списка.
type Node struct {
	data interface{}
	next *Node
}

// NewDynArr создает и возвращает новый экземпляр DynArr.
func NewDynArr() *DynArr {
	return &DynArr{}
}

// Add добавляет новый элемент в конец динамического массива.
func (da *DynArr) Add(data interface{}) {
	newNode := &Node{data: data}

	if da.head == nil {
		da.head = newNode
	} else {
		current := da.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

// Remove удаляет указанный элемент из динамического массива.
func (da *DynArr) Remove(data interface{}) {
	if da.head == nil {
		return
	}

	if da.head.data == data {
		da.head = da.head.next
		return
	}

	current := da.head
	for current.next != nil {
		if current.next.data == data {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

// BinarySearch выполняет двоичный поиск указанного элемента в динамическом массиве.
// Возвращает индекс элемента, если он найден, и -1, если не найден.
func (da *DynArr) BinarySearch(target interface{}) int {
	index := 0
	current := da.head
	for current != nil {
		if current.data == target {
			return index
		}
		current = current.next
		index++
	}
	return -1
}

func main() {
	dynArr := NewDynArr()
	dynArr.Add(10)
	dynArr.Add(20)
	dynArr.Add(30)

	fmt.Printf("Данные перед удалением: ")
	dynArr.Print()
	dynArr.Remove(20)
	fmt.Printf("Данные после удаления: ")
	dynArr.Print()

	index := dynArr.BinarySearch(10)
	if index != -1 {
		fmt.Printf("Элемент 10 найден по индексу %d\n", index)
	} else {
		fmt.Println("Элемент 10 не найден")
	}
}

// Print выводит элементы динамического массива.
func (da *DynArr) Print() {
	current := da.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}

`
Шаг 1: Создание экземпляра NewDynArr
   dynArr
   +------+
   | head |
   +------+

Шаг 2: Добавление элементов 10, 20 и 30
   dynArr
   +------+    +---+    +---+    +---+
   | head | -> | 10| -> | 20| -> | 30|
   +------+    +---+    +---+    +---+

Шаг 3: Вывод данных перед удалением
   Данные перед удалением: 10 20 30

Шаг 4: Удаление элемента 20
   dynArr
   +------+    +---+         +---+
   | head | -> | 10|   ->    | 30|
   +------+    +---+         +---+

Шаг 5: Вывод данных после удаления
   Данные после удаления: 10 30

Шаг 6: Двоичный поиск элемента 10
   Результат: Элемент 10 найден по индексу 0

`