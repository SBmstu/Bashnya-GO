package utils

import "fmt"

func PrintResult(arr []string) {
	fmt.Println("Результат: ");
	for i, s := range arr {
		fmt.Printf("%d: %s\n", i, s);
	}
}
