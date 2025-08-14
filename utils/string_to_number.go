// Package utils serve para funções auxiliares
package utils

import "strconv"


// StringToNumbers recebe um slice de strings, processa e retorna um slice do tipo float64
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
