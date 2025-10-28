package main

// Попробовать с помощью жэнэриков сделать универсальный стек, если так можно

import (
	"bashnya-hw3/stack"
	"bashnya-hw3/util"
	"fmt"
	"os"
)

func main() {
	var err error = run();
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "\nОшибка: %v\n", err); // разобраться со спецификаторами
		os.Exit(1);
	}

	os.Exit(0); // Как лучше обрабатывать ошибки?
}

func run() error {
	var stack stack.Stack_int; // не будет ругаться на имя переменной?
	stack.Init();

	var err error;
	var choice int = -1;
	for (choice != 0) {
		err = util.ChooseOption(&choice);
		if (err != nil) {
			return err;
		}

		switch choice {
		case util.EXIT:
			return nil;

		case util.PUSH:
			err = util.StackPush(&stack)
			if (err != nil) {
				return err;
			}

		case util.POP:
			util.StackPop(&stack)

		case util.SIZE:
			util.StackSize(&stack)

		case util.CLEAR:
			stack.Clear()

		// case util.ADD:
		// 	return err
		}
	}

	return nil;
}
