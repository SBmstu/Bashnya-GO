package stack

import (
	"fmt"
	"unsafe"
)

var empty int;
const STACK_TYPE_SIZE = unsafe.Sizeof(empty);
const STACK_SIZE = 1024 * 1024;
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
	top int // указатель стека
}

func (s *Stack_int) Init() {
	s.stack_array = make([]int, 0, STACK_SIZE);
	s.top = 0;

	fmt.Println("Стек успешно инициализирован");
}

func (s *Stack_int) Pop(n int) error {
	s.stack_array[s.top] = n
	s.top++

	return nil
}

func (s *Stack_int) Push(n int) error {

	// fmt.Println("Элемент успешно добавлен в стек")
	s.stack_array[s.top] = n;
	s.top++

	return nil
}

func (s *Stack_int) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack_int) Size() int {
	return s.top
}

func (s *Stack_int) Clear() {
	s.stack_array = make([]int, 0, STACK_SIZE) // Так ли?
	s.top = 0

	// fmt.Println("Стек успешно очищен")
}
