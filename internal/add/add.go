// Package add do addition math
package add

import "errors"

const err = "error: not enough numbers"

func AddInteger(numbers ...int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New(err)
	}
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total, nil
}

func AddFloat(numbers ...float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, errors.New(err)
	}
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total, nil
}
