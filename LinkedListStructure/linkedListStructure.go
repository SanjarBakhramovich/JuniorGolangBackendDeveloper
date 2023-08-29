package main

// Определение структуры для узла двусвязного списка
type Node struct {
	Prev *Node
	Next *Node
	Data interface{}
}

func NewNode(prev *Node, next *Node, data interface{}) *Node {
	var node Node
	node.Prev = prev
	node.Next = next
	node.Data = data

	return &node
}

func main() {
	// var n1 Node
	// var n2 Node
	// val := 1
	// node1 := Constructor(&n1, &n2, val)

	// node2 := Node{
	// 	Prev: &n1,
	// 	Next: &n2,
	// 	Data: val,
	// }

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

// Метод для удаления узла с определенными данными из списка
func (list *DoublyLinkedList) Remove(data interface{}) {
	nodeToRemove := list.Find(data)
	if nodeToRemove == nil {
		return
	}

	// Обновляем связи соседних узлов после удаления узла
	if nodeToRemove.Prev != nil {
		nodeToRemove.Prev.Next = nodeToRemove.Next
	} else {
		list.Head = nodeToRemove.Next
	}

	if nodeToRemove.Next != nil {
		nodeToRemove.Next.Prev = nodeToRemove.Prev
	} else {
		list.Tail = nodeToRemove.Prev
	}
}

// func main() {
// 	// Создание нового двусвязного списка
// 	linkedList := NewDoublyLinkedList()
// 	linkedList.Add(1)
// 	linkedList.Add(2)
// 	linkedList.Add(3)

// 	// Поиск узла с данными 2 в списке
// 	foundNode := linkedList.Find(2)
// 	if foundNode != nil {
// 		fmt.Println(foundNode.Data)
// 	} else {
// 		fmt.Println("Node not found")
// 	}

// 	// Удаление узла с данными 2 из списка
// 	linkedList.Remove(2)
// 	foundNode = linkedList.Find(2)
// 	if foundNode != nil {
// 		fmt.Println(foundNode.Data)
// 	} else {
// 		fmt.Println("Node not found")
// 	}
// }
