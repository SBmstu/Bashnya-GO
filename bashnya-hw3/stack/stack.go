package stack

import "fmt"

const STACK_SIZE = 1024 * 1024;

type Stack interface {
	Pop()
	Push()
	IsEmpty()
	Size()
	Clear()
}

type Stack_int struct {
	stack_array []int
	top int
}

func (s *Stack_int) Init() {
	s.stack_array = make([]int, 0, STACK_SIZE);
	s.top = 0;

	fmt.Println("Стек успешно инициализирован");
}

func (s *Stack_int) Pop() {

}

func (s *Stack_int) Push() {

}

func (s *Stack_int) IsEmpty() {

}

func (s *Stack_int) Size() {

}

func (s *Stack_int) Clear() {

}
