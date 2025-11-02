package utils

import (
	"bufio"
	"fmt"
)

// func GetReader(cfg *Config_t) (io.Reader, error) {
// 	if (cfg.inputFile == "") {
// 		return os.Stdin, nil
// 	}

// 	file, err := os.Open(cfg.inputFile)
// 	if (err != nil) {
// 		return nil, err
// 	}

// 	return file, nil
// }

func ReadData(cfg *Config_t) ([]string, error) {	
	var data []string;

	fmt.Println("Введите строки: ");
	scanner := bufio.NewScanner(cfg.inputStream);
	for scanner.Scan() {
		line := scanner.Text();
		data = append(data, line);
	}

	return data, nil;
}

func GenMapWithCount(data []string) map[string]int {
	res := make(map[string]int);

	for _, s := range data {
		res[s]++;
	}

	return res;
}
