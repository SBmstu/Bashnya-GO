package main

import (
	"bashnya-hw4/utils"
	"fmt"
	"os"
)

func main() {
	err := run();
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "\nОшибка: %v\n", err);
		os.Exit(1);
	}

	os.Exit(0);
}

func run() error {
	key, err := utils.ProcessArgs(os.Args);
	if (err != nil) {
		return err;
	}

	switch (key) {
	case 1:
		fmt.Println("One");
	default:
		fmt.Println("Hz");
	}

	return nil;
}
