// Package sub return a minus b
package sub

import "errors"

const err = "error: not enough numbers"

func SubInteger(numbers ...int) (int, error) {
	if len(numbers) < 2 {
		return 0, errors.New(err)
	}
	total := 0
	for _, number := range numbers {
		total -= number
	}
	return total, nil
}

func SubFloat(numbers ...float64) (float64, error) {
	if len(numbers) < 2 {
		return 0.0, errors.New(err)
	}
	total := 0.0
	for _, number := range numbers {
		total -= number
	}
	return total, nil
}
