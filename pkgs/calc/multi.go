package calc

import (
	"errors"
	"fmt"
	"math"
)

// Multi recebe uma quantidade variada de valores do tipo float64 e retorna uma string e um erro
func Multi(values ...float64) (string, error) {
	if len(values) == 0 {
		return "error:", errors.New(err)
	}

	total := values[0]
	for _, value := range values[1:] {
		total *= value
	}

	if total == math.Trunc(total) {
		return fmt.Sprintf("%d", int64(total)), nil
	}
	return fmt.Sprintf("%g", total), nil
}
