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
		Short: "CLI Tool for simple math",
		Long: "CLI Tool build with cobra and go for simple math",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to GoCalc")
			fmt.Println("Use gocalc -h to see the commands")
		},
	}

	var cmdAdd = &cobra.Command{
		Use: "add",
		Short: "add n numbers",
		Long: "receives a variable quantity of numbers and returns the sum of them",
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
		Short: "sub n numbers",
		Long: "receives a variablle quantity of numbers and returns the sub of them",
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
		Short: "multiply n numbers",
		Long: "receives a variable quatity of numbers and returns the sub of them",
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
