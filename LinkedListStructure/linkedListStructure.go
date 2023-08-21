package main

import "fmt"

type Node struct {
	Prev *Node
	Next *Node
	Data interface{}
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (list *DoublyLinkedList) Add(data interface{}) {
	newNode := &Node{Data: data}

	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
	} else {
		newNode.Prev = list.Tail
		list.Tail.Next = newNode
		list.Tail = newNode
	}
}

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

func (list *DoublyLinkedList) Remove(data interface{}) {
	nodeToRemove := list.Find(data)
	if nodeToRemove == nil {
		return
	}

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

func main() {
	linkedList := NewDoublyLinkedList()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)

	foundNode := linkedList.Find(2)
	if foundNode != nil {
		fmt.Println(foundNode.Data)
	} else {
		fmt.Println("Node not found")
	}

	linkedList.Remove(2)
	foundNode = linkedList.Find(2)
	if foundNode != nil {
		fmt.Println(foundNode.Data)
	} else {
		fmt.Println("Node not found")
	}
}
