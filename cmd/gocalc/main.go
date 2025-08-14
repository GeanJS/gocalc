package main

import (
	"fmt"
	"gocalc/pkgs/calc"
	"gocalc/utils"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "gocalc",
		Short: "Ferramenta CLI que realiza operações matemáticas básicas",
		Long: "Ferramenta CLI escrita em Go usando o pacote Cobra, para realizar operações matemáticas básicas",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Bem-vindo ao GoCalc")
			fmt.Println("Use o comando help para ver as operações disponiveis")
		},
	}

	var cmdAdd = &cobra.Command{
		Use: "add",
		Short: "soma n numeros",
		Long: "recebe uma quantidade variavel de numeros e retorna o resultado da adição",
		Run: func(cmd *cobra.Command, args []string) {
			numbers, err := utils.StringToNumbers(args)
			if err != nil {
				return 
			}
			result, _ := calc.Add(numbers...)
			fmt.Println("Result: ", result)
		},
	}
	var cmdSub = &cobra.Command {
		Use: "sub",
		Short: "subtrai n numeros",
		Long: "recebe uma quantidade variavel de numeros e retorna o resultado da subtração",
		Run: func(cmd *cobra.Command, args []string) {
			numbers, err := utils.StringToNumbers(args)
			if err != nil {
				return
			}

			result, _ := calc.Sub(numbers...)
			fmt.Println("Result: ", result)
		},
	}
	var cmdMulti = &cobra.Command{
		Use: "multi",
		Short: "multiplica n numeros",
		Long: "recebe uma quantidade variavel de numeros e retorna o resultado da multiplicação",
		Run: func(cmd *cobra.Command, args []string) {
			numbers, err := utils.StringToNumbers(args)
			if err != nil {
				return
			}

			result, _ := calc.Multi(numbers...)
			fmt.Println("Result: ", result)
		},
	}
	var cmdDiv = &cobra.Command{
		Use: "div",
		Short: "divide n numeros",
		Long: "recebe uma quantidade variavel de numeros e retorna o resultado da divisão",
		Run: func(cmd *cobra.Command, args []string) {
			numbers, err := utils.StringToNumbers(args)
			if err != nil {
				return
			}

			result, _ := calc.Div(numbers...)
			fmt.Println("Result: ", result)
		},
	}

	rootCmd.AddCommand(cmdDiv)
	rootCmd.AddCommand(cmdMulti)
	rootCmd.AddCommand(cmdSub)
	rootCmd.AddCommand(cmdAdd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
}
