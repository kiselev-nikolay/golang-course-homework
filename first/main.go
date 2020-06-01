package main

import (
	"fmt"
	"os"
	"strconv"
)

const dollarExcangeRate float64 = 69.3

func main() {
	if len(os.Args) <= 1 {
		printHelp()
	}
	arg := os.Args[1]
	switch arg {
	case "--help":
		printHelp()
	case "exchange":
		rubles, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Println("Error with argument rubles. Must be a float.")
			os.Exit(1)
		}
		result := RubleToDollarExcangeCalculator(rubles)
		fmt.Printf("%.2f$\n", result)
	}
}

func printHelp() {
	fmt.Println("Super strange calc 1.0")
	fmt.Println("commands: excange, geometry, debit")
}

// Программа для конвертации рублей в доллары. Программа запрашивает сумму в рублях и выводит сумму в долларах.
// Курс доллара задайте константой.
func RubleToDollarExcangeCalculator(rubles float64) float64 {
	return rubles / dollarExcangeRate
}
