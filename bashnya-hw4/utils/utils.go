package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func GetReader(cfg *Config_t) (io.Reader, error) {
	if (cfg.inputFile == "") {
		return os.Stdin, nil
	}

	file, err := os.Open(cfg.inputFile)
	if (err != nil) {
		return nil, err
	}
	
	return file, nil
}

func ReadData(reader io.Reader) ([]string, error) {	
	var data []string;

	fmt.Println("Введите строки: ");
	scanner := bufio.NewScanner(reader);
	for scanner.Scan() {
		line := scanner.Text();
		data = append(data, line);
	}

	return data, nil;
}
