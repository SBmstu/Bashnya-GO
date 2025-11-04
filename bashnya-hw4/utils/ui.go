package utils

import (
	"fmt"
)

func PrintResult(arr []string, cfg *Config_t) {
	// writer := bufio.NewWriter(cfg.OutputStream);
	
	fmt.Println("Результат: ");
	for _, s := range arr {
		fmt.Printf("%s\n", s);
	}
}
