package utils

import (
	"bufio"
	"fmt"
	"strings"
)

type Data_t struct {
	Original []string
	Configured []string
	Count int
}


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

func ConfigurateString(s string, cfg *Config_t) string {
	if (cfg.i) {
		s = strings.ToLower(s);
	}

	if (cfg.numFields != 0 && cfg.numChars == 0) {
		s = strings.Join(strings.Split(s, " ")[cfg.numFields:], " ");
	}
	
	if (cfg.numChars != 0 && cfg.numFields == 0) {
		s = s[cfg.numChars:]; // А если cfg.numChars > len(s) ?
	} else if (cfg.numChars != 0 && cfg.numFields != 0) {
		// +1 или нет?
		s = s[:cfg.numChars + 1]; // А если cfg.numChars > len(s) ?
	}

	return s;
}

func ConfigurateData(data *Data_t, cfg *Config_t) {
	configurated := make([]string, data.Count);

	for i, s := range data.Original {
		configurated[i] = ConfigurateString(s, cfg);
	}

	data.Configured = configurated;
}

func GenMapWithCount(data *Data_t, cfg *Config_t) map[string]int {
	res := make(map[string]int);

	for _, s := range data.Original {
		res[strings.ToLower(s)]++;
	}

	return res;
}
