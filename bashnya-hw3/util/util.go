package util

import (
	"errors"
	"fmt"
)

var (
	ErrorInput = errors.New("input error")
	ErrorRange = errors.New("range error")
)

func ChooseOption(choice *int) error {
	PrintOptions();
	fmt.Printf("Выберите опцию: ")
	_, err := fmt.Scan(choice);
	if (err != nil) {
		return ErrorInput
	}

	if (*choice < 0 || *choice > OPTINOS_COUNT) {
		return ErrorRange
	}

	return nil;
}
