package main

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

			return err;

		case util.POP:
			return err;
		case util.SIZE:
			return err;

		case util.CLEAR:
			return err;

		case util.ADD:
			return err
		}
	}

	return nil;
}
