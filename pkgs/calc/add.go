// Package calc has the funcs of a basic calculator
package calc

import (
	"errors"
	"fmt"
	"math"
)

const err = "not enough numbers"

func Add(values ...float64) (string, error) {
	if len(values) == 0 {
		return "Error", errors.New(err)
	}

	var total float64
	for _, value := range values {
		total += value
	}

	if total == math.Trunc(total) {
		return fmt.Sprintf("%d", int64(total)), nil
	}
	return fmt.Sprintf("%g", total), nil
}
