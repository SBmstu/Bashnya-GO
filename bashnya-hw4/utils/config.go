package utils

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	ErrorInput = errors.New("input error");
	ErrorWrongArgsNumber = errors.New("wrong number of arguments");
)

type Config interface {
	Init()
	Print()
}

type Config_t struct {
	c bool
	d bool
	u bool
	i bool 
	numFields int
	numChars int
	inputStream io.Reader
	outputStream io.Writer
}

func (cfg *Config_t) GenConfig() error {
	flag.BoolVar(&cfg.c, "c", false, "подсчет количества встреч строки во входных данных");
	flag.BoolVar(&cfg.d, "d", false, "вывод только тех строк, которые повторяются")
	flag.BoolVar(&cfg.u, "u", false, "вывод только тех строк, которые не повторяются во входных данных");
	flag.BoolVar(&cfg.i, "i", false, "не учитывать регистр букв");
	flag.IntVar(&cfg.numFields, "f", 0, "не учитывать первые num_fields полей в строке. Полем в строке является непустой набор символов отделенным пробелом");
	flag.IntVar(&cfg.numChars, "s", 0, "не учитывать первые num_chars символов в строке. При использовании вместе с параметром -f учитываются первые символы после num_fields полей (не учитывая пробел-разделитель после последнего поля)");

	flag.Parse();

	args := flag.Args()
	switch len(args) {
	case 0:
		cfg.inputStream = os.Stdin
		cfg.outputStream = os.Stdout
	case 1:
		file, err := os.Open(args[0])
		if (err != nil) {
			return err
		}
		cfg.inputStream = file;
		cfg.outputStream = os.Stdout
	case 2:
		file_input, err := os.Open(args[0])
		if (err != nil) {
			return err
		}
		file_output, err := os.Open(args[1])
		if (err != nil) {
			return err
		}

		cfg.inputStream = file_input;
		cfg.outputStream = file_output;
	default:
		return ErrorWrongArgsNumber
	}

	return nil;
}

func (cfg *Config_t) GetFlags() map[string]bool {
	flags := map[string]bool {
		"c": cfg.c,
		"d": cfg.d,
		"u": cfg.u,
		"i": cfg.i,
	}

	return flags
}

func (cfg *Config_t) Print() {
	fmt.Printf("\nActive flags: ", );
	for flag, activity := range cfg.GetFlags() {
		if (activity) {
			fmt.Printf("%s ", flag);
		}
	}

	fmt.Println();

	fmt.Printf("Unactive flags: ");
	for flag, activity := range cfg.GetFlags() {
		if (!activity) {
			fmt.Printf("%s ", flag);
		}
	}

	fmt.Println();

	fmt.Printf("numField: %d; numChar: %d\n", cfg.numFields, cfg.numChars);
	// fmt.Printf("inputStream: %s; outputStream: %s\n\n", cfg.inputStream, cfg.outputStream); // Подумать, как тут сделать
}
