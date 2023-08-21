package main

import "fmt"

type Slice struct {
	data []interface{}
}

func NewSlice() *Slice {
	return &Slice{}
}

func (s *Slice) Add(data interface{}) {
	s.data = append(s.data, data)
}

func (s *Slice) Remove(data interface{}) {
	for i, val := range s.data {
		if val == data {
			s.data = append(s.data[:i], s.data[i+1:]...)
			return
		}
	}
}

func main() {
	slice := NewSlice()
	slice.Add(10)
	slice.Add(20)
	slice.Add(30)

	fmt.Println(slice.data)
	slice.Remove(20)
	fmt.Println(slice.data)
}
