// Package utils
package utils

import "strconv"


func StringToNumbers(args []string) ([]float64, error) {
	numbers := make([]float64, len(args))

	for i, s := range args {
		num, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, err
		}
		numbers[i] = num

	}
	return numbers, nil
	
}
