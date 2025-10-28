package stack

import (
	"fmt"
	"unsafe"
)

var empty int;
const STACK_TYPE_SIZE = int(unsafe.Sizeof(empty));
const STACK_SIZE int = 1024 * 1024;
const STACK_MAX_ELEMS_COUNT = STACK_SIZE / STACK_TYPE_SIZE;

type Stack interface {
	Pop()
	Push()
	IsEmpty()
	Size()
	Clear()
}

type Stack_int struct {
	stack_array []int // массив стека
	stack_size int // количество свободных ячеек
	top int // указатель стека
}

func (s *Stack_int) Init() {
	s.stack_array = make([]int, 0, STACK_SIZE);
	s.stack_size = STACK_MAX_ELEMS_COUNT;
	s.top = 0; // Может с -1?

	fmt.Println("Стек успешно инициализирован");
}

func (s *Stack_int) Pop() (int, bool) {
	var index int = s.top;
	if (index == 0) {
		return -1, false
	}

	var elem int = s.stack_array[index - 1];
	s.stack_array = s.stack_array[:index - 1];
	s.top--;
	s.stack_size++;

	return elem, true
}

// Правильно ли я понял значение этого метода?
// Или нужно было через append()?
func (s *Stack_int) Push(n int) error {
	s.stack_array = append(s.stack_array, n)

	s.top++
	s.stack_size--;

	return nil
}

func (s *Stack_int) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack_int) Size() (int, int) {
	return s.top, s.stack_size
}

func (s *Stack_int) Clear() {
	s.stack_array = make([]int, 0, STACK_SIZE) // Так ли? Или эффективнее пройтись по списку и очистить его 
	s.top = 0

	fmt.Println("Стек успешно очищен")
}
