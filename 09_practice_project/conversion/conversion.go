package conversion

import (
	"errors"

	"strconv"
)

func StringToFloat(strings []string) ([]float64, error) {
	var prices []float64
	for _, line := range strings {
		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return nil, errors.New("Failed to parse")
		}
		prices = append(prices, floatPrice)
	}
	return prices, nil
}
