package utils

import (
	"io"
	"os"
)

func ProcessFlags(cfg *Config_t) {
	if (cfg.c) {

	} else if (cfg.d) {

	} else if (cfg.u) {

	}
}

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
