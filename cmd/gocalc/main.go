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

	rootCmd.AddCommand(cmdAdd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	
}
