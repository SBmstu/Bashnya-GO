package utils

import "fmt"

func PrintResult(arr []string, cfg *Config_t) {
	// writer := 
	fmt.Println("Результат: ");
	for i, s := range arr {
		fmt.Printf("%d: %s\n", i, s);
	}
}
