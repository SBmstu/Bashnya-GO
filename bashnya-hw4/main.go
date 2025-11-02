package main

import (
	"bashnya-hw4/processing"
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
	var cfg utils.Config_t;
	err := cfg.GenConfig();
	if (err != nil) {
		return err;
	}
	
	data_arr, err := utils.ReadData(&cfg); // Почему я не могу здесь обращаться к полям структуры cfg?
	if (err != nil) {
		return err;
	}

	var data utils.Data_t;
	data.Original = data_arr;
	data.Count = len(data_arr);

	result := processing.ProcessData(&data, &cfg);
	utils.PrintResult(result, &cfg);

	return nil;
}
