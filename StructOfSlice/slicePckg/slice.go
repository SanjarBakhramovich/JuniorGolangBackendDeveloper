package slicePckg

const ArrCap = 4

type StructOfArr struct {
	Length int
	Cap    int
	Arr    *[ArrCap]int
	// Интерфейс вроде можно
	// Вернуть новую структура newStructArr
}

func BinSearch(arr *[ArrCap]int, Length, Target int) int {
	low := 0
	high := Length - 1

	for low <= high {
		mid := low + (high-low)/2 //идёт в центр

		if arr[mid] == Target {
			return mid // если нашёл возвращает
		} else if arr[mid] < Target {
			low = mid + 1 // идёт вправо
		} else {
			high = mid - 1 // идёт влево
		}
	}
	// выход из лупа
	return -1 //если цифра не найдена, возвращает -1 ошибку
}

func (s *StructOfArr) Sort() {
	for i := 0; i < s.Length-1; i++ {
		for j := 0; j < s.Length-i-1; j++ {
			if s.Arr[j] > s.Arr[j+1] {
				s.Arr[j], s.Arr[j+1] = s.Arr[j+1], s.Arr[j]
			}
		}
	}
}

// CRUD
// 1
func (s *StructOfArr) Add(num int) bool {
	if s.Length >= s.Cap {
		return false
	}

	s.Arr[s.Length] = num
	s.Length++
	return true
}

// 2
func (s *StructOfArr) DeleteVal(value int) bool {
	for i := 0; i < s.Length; i++ {
		if s.Arr[i] == value {
			for j := i; j < s.Length-1; j++ {
				s.Arr[j] = s.Arr[j+1]
			}
			s.Length--
			return true
		}
	}
	return false
}

// 3
func (s *StructOfArr) GetSlice() []int {
	slice := s.Arr[:s.Length]
	for i := 0; i < s.Length; i++ {
		slice[i] = s.Arr[i]
	}
	return slice
}

// 4
func (s *StructOfArr) Update(index, newVal int) bool {
	if index < 0 || index >= s.Length {
		return false
	}

	if s.Length >= s.Cap {
		newCap := s.Cap * 2
		newArr := [ArrCap]int{}

		for i := 0; i < s.Length; i++ {
			newArr[i] = s.Arr[i]
		}

		s.Arr = &newArr
		s.Cap = newCap
	}

	s.Arr[index] = newVal
	return true
}

// нужна логика если массив заполнен то нужно
// создать новый конструктор который умножит на 2 массив
// и вернёт скопированный вариант
//
