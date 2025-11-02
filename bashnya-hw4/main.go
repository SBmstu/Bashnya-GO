package main

import (
	"bashnya-hw4/functions"
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

	reader, err := utils.GetReader(&cfg)
	if (err != nil) {
		return err;
	}

	data, err := utils.ReadData(reader);
	if (err != nil) {
		return err;
	}

	result := functions.FindUniq(data);
	utils.PrintResult(result);

	return nil;
}
