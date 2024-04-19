package conversion

import (
	"errors"
	"strconv"
)

//package to conver to float64/string and int conversions

// receives slice string nad return a float
func StringToFloats(strings []string) ([]float64, error) {

	var floats []float64
	for _, stringVal := range strings {
		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("Failed to parse float")
		}

		floats = append(floats, floatVal)

	}
	return floats, nil
}
