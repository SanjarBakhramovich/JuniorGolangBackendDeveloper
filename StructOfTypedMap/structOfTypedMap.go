package main

import "fmt"

type IntMap struct {
	data map[string]int
}

func NewIntMap() *IntMap {
	return &IntMap{
		data: make(map[string]int),
	}
}

func (m *IntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *IntMap) Get(key string) (int, bool) {
	value, exists := m.data[key]
	return value, exists
}

func (m *IntMap) Remove(key string) {
	delete(m.data, key)
}

func main() {
	intMap := NewIntMap()
	intMap.Add("one", 1)
	intMap.Add("two", 2)
	intMap.Add("three", 3)

	fmt.Println(intMap.Get("two"))
	intMap.Remove("two")
	fmt.Println(intMap.Get("two"))
}

//Написать Struct of slice without using slice
// Struct should include cap, len, pointer to arr
// cap of array should be assigned by const
// major methods should be add, delete, and binSearch [CRUD]
