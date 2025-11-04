package main

import (
	"bashnya-hw4/processing"
	"bashnya-hw4/utils"
	"fmt"
	"os"
)

// есть момент, что выводится не в порядке ввода
// + в output пустая строка заносится

func main() {
	err := run();
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "\nОшибка: %v\n", err);
		os.Exit(1);
	}

	os.Exit(0);
}

func run() error {
	var cfg utils.Config_t;
	err := cfg.GenConfig();
	if (err != nil) {
		return err;
	}

	data_arr, err := utils.ReadData(&cfg);
	if (err != nil) {
		return err;
	} // а что делать если в InputStream файл? Как закрывать то? Тот же самый вопрос и с OutputStream

	var data utils.Data_t;
	data.Original = data_arr;
	data.Count = len(data_arr);

	result := processing.ProcessData(&data, &cfg);
	
	err = utils.PrintResult(result, &cfg);
	if (err != nil) {
		return err;
	}

	return nil;
}
