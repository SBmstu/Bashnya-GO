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

type ResInfo_t struct {
	I int
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
	scanner := bufio.NewScanner(cfg.InputStream);
	for scanner.Scan() {
		line := scanner.Text();
		data = append(data, line);
	}

	return data, nil;
}

func ConfigurateString(s string, cfg *Config_t) string {
	if (cfg.I) {
		s = strings.ToLower(s);
	}

	if (cfg.NumFields != 0 && cfg.NumChars == 0) {
		s = strings.Join(strings.Split(s, " ")[cfg.NumFields:], " ");
	}
	
	if (cfg.NumChars != 0 && cfg.NumFields == 0) {
		s = s[cfg.NumChars:]; // А если cfg.numChars > len(s) ?
	} else if (cfg.NumChars != 0 && cfg.NumFields != 0) {
		// +1 или нет?
		s = s[:cfg.NumChars + 1]; // А если cfg.numChars > len(s) ?
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

func GenMapWithCount(data *Data_t, cfg *Config_t) map[string]ResInfo_t {
	res := make(map[string]ResInfo_t);

	ConfigurateData(data, cfg);

	for i, s := range data.Configured {
		count := res[s].Count + 1;
		res[s] = ResInfo_t{i, count}; // почему только так получается?
	}

	return res;
}
