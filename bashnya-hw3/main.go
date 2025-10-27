package main

import (
	"bashnya-hw3/stack"
	"bashnya-hw3/util"
)

func main() {
	var stack stack.Stack_int; // не будет ругаться на имя переменной?
	stack.Init();

	var choice int = -1;
	for (choice != 0) {
		err := util.ChooseOption(&choice);
		if (err != nil) {
			return ;
		}

		switch choice {
		case util.EXIT:
			return
		case util.PUSH:
			return
		case util.POP:
			return
		case util.SIZE:
			return
		case util.CLEAR:
			return
		case util.ADD:
			return
		}
	}
}
