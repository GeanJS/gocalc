package calc

import (
	"errors"
	"fmt"
	"math"
)

// Div recebe uma quantidade variada de valores do tipo float64, faz a divisão desses valores e retorna uma string ou erro
func Div(values ...float64) (string, error) {
	if len(values) == 0 {
		return "error: ", errors.New(err)
	}

	total := values[0]
	for _, value := range values[1:] {
		total /= value
	}

	if total == math.Trunc(total) {
		return fmt.Sprintf("%d", int64(total)), nil
	}

	return fmt.Sprintf("%g", total), nil
}
