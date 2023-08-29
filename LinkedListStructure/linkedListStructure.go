package main

import "fmt"

// Определение структуры для узла двусвязного списка
type Node struct {
	Prev *Node
	Next *Node
	Data interface{}
}

// NewNode создает и возвращает новый узел двусвязного списка с указанными данными.
// Параметры:
//
//	prev: Указатель на предыдущий узел.
//	next: Указатель на следующий узел.
//	data: Данные, которые будут храниться в узле.
//
// Возвращает:
//
//	Указатель на созданный узел.
func NewNode(prev *Node, next *Node, data interface{}) *Node {
	var node Node
	node.Prev = prev
	node.Next = next
	node.Data = data

	return &node
}

// Определение структуры для двусвязного списка
type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

// Функция для создания нового пустого двусвязного списка
// конструктор
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

// Метод для добавления нового узла с данными в конец списка
func (list *DoublyLinkedList) Add(data interface{}) {
	newNode := &Node{Data: data}

	// Если список пустой, устанавливаем новый узел как голову и хвост
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		// Если список не пустой, добавляем новый узел после хвостового узла
		newNode.Prev = list.Tail
		list.Tail.Next = newNode
		list.Tail = newNode
	}
}

// Метод для поиска узла по данным в списке
func (list *DoublyLinkedList) Find(data interface{}) *Node {
	currentNode := list.Head
	for currentNode != nil {
		if currentNode.Data == data {
			return currentNode
		}
		currentNode = currentNode.Next
	}
	return nil
}

// Remove удаляет первый узел с указанными данными из двусвязного списка, если он существует.
// Если узел с указанными данными не найден, ничего не происходит.
// Параметры:
//
//	data: Данные, по которым будет производиться поиск и удаление узла.
func (list *DoublyLinkedList) Remove(data interface{}) {
	// Находим узел для удаления
	nodeToRemove := list.Find(data)

	// Если узел не найден, завершаем выполнение
	if nodeToRemove == nil {
		return
	}

	// Обновляем связи соседних узлов после удаления узла
	if nodeToRemove.Prev != nil {
		nodeToRemove.Prev.Next = nodeToRemove.Next
	} else {
		// Если удаляется головной узел, обновляем голову списка
		list.Head = nodeToRemove.Next
	}

	if nodeToRemove.Next != nil {
		nodeToRemove.Next.Prev = nodeToRemove.Prev
	} else {
		// Если удаляется хвостовой узел, обновляем хвост списка
		list.Tail = nodeToRemove.Prev
	}
}

func main() {
	// Создание нового двусвязного списка
	linkedList := NewDoublyLinkedList()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)

	// Поиск узла с данными 2 в списке
	foundNode := linkedList.Find(2)
	if foundNode != nil {
		fmt.Println(foundNode.Data)
	} else {
		fmt.Println("Node not found")
	}

	// Удаление узла с данными 2 из списка
	linkedList.Remove(2)
	foundNode = linkedList.Find(2)
	if foundNode != nil {
		fmt.Println(foundNode.Data)
	} else {
		fmt.Println("Node not found")
	}
}

`
Шаг 1: Создание нового пустого двусвязного списка
   linkedList
   +-------+
   |  Head | ---> nil
   +-------+
   |  Tail | ---> nil
   +-------+

Шаг 2: Добавление элемента 1 в список
   linkedList
   +-------+
   |  Head | ---> Node1 (1)
   +-------+
   |  Tail | ---> Node1 (1)
   +-------+

Шаг 3: Добавление элемента 2 в список
   linkedList
   +-------+
   |  Head | ---> Node1 (1)
   +-------+
   |  Tail | ---> Node2 (2)
   +-------+

Шаг 4: Добавление элемента 3 в список
   linkedList
   +-------+
   |  Head | ---> Node1 (1)
   +-------+
   |  Tail | ---> Node3 (3)
   +-------+

Шаг 5: Поиск узла с данными 2 в списке
   Результат: Найден узел с данными 2

Шаг 6: Удаление узла с данными 2 из списка
   linkedList
   +-------+
   |  Head | ---> Node1 (1)
   +-------+
   |  Tail | ---> Node3 (3)
   +-------+

Шаг 7: Поиск узла с данными 2 в списке после удаления
   Результат: Узел не найден

`