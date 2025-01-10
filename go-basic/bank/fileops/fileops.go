package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadDataFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to read")
	}
	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 1000, errors.New("failed to reade")
	}
	return balance, nil
}

func WriteDataToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
