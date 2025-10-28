package util

import (
	"bashnya-hw3/stack"
	"errors"
	"fmt"
)

var (
	ErrorInput = errors.New("input error")
	ErrorRange = errors.New("range error")
)

func ChooseOption(choice *int) error {
	PrintOptions();
	fmt.Printf("Выберите опцию: ")
	_, err := fmt.Scan(choice);
	if (err != nil) {
		return ErrorInput
	}

	if (*choice < 0 || *choice > OPTINOS_COUNT) {
		return ErrorRange
	}

	fmt.Println();

	return nil;
}

func StackPop(s *stack.Stack_int) {
	elem, not_empty := s.Pop()
	if (!not_empty) {
		fmt.Println("Стек пуст")
		return
	}

	fmt.Printf("Удаленный элемент: %d\n", elem);
}

func StackPush(s *stack.Stack_int) error {
	var n int
	fmt.Printf("Введите число, которое хотели бы внести в стек: ");
	_, err := fmt.Scan(&n);
	if (err != nil) {
		return ErrorInput
	}

	err = s.Push(n);
	if (err != nil) {
		return err
	}

	fmt.Printf("Элемент успешно добавлен в стек\n\n")

	return nil;
}

func StackSize(s *stack.Stack_int) {
	elems_count, free_cells := s.Size()

	fmt.Printf("Количество элементов: %d\nКоличество свободных ячеек: %d", elems_count, free_cells);
}
