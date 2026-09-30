package utils

import (
	"bufio"
	"errors"
)

var (
	ErrorBufferFlush error = errors.New("ошибка сброса буфера")
	ErrorStreamWrite error = errors.New("ошибка записи в поток")
)

func PrintResult(arr []string, cfg *Config_t) error {
	writer := bufio.NewWriter(cfg.OutputStream);
	
	// fmt.Println("Результат: ");
	for _, s := range arr {
		_, err := writer.WriteString(s + "\n");
		if (err != nil) {
			return ErrorStreamWrite;
		}
	}

	err := writer.Flush();
    if err != nil {
        return ErrorBufferFlush;
    }

	return nil;
}
