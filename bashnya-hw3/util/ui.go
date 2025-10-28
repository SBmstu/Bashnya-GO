package util

import "fmt"

const OPTINOS_COUNT = 5;

const (
	EXIT = iota
	PUSH
	POP
	SIZE
	CLEAR
	// ADD
)

var options = [OPTINOS_COUNT]string{
	"Завершить программу",
	"Добавить элемент в стек",
	"Удалить элемент из стека",
	"Получить размер стека",
	"Очистить стек",
};

func PrintOptions() {
	for i, option := range options {
		fmt.Printf("%d - %s\n", i, option);
	}
}
