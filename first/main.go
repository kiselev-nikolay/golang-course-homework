package main

import (
	"fmt"
	"math"
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
	case "geometry":
		sideA, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Println("Error with sideA. Must be a float.")
			os.Exit(1)
		}
		sideB, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Println("Error with sideB. Must be a float.")
			os.Exit(1)
		}
		area, perimeter, sideC := GeometryCalc(sideA, sideB)
		fmt.Printf("area=%.4f perimeter=%.4f sideC=%.4f\n", area, perimeter, sideC)
	case "debit":
		debit, err := strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			fmt.Println("Error with debit. Must be a float.")
			os.Exit(1)
		}
		profitPercent, err := strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Println("Error with profit percent. Must be a float.")
			os.Exit(1)
		}
		result := DebitCalc(debit, profitPercent)
		fmt.Printf("%.4f with %.4f%% for 5 years equal %.4f\n", debit, profitPercent, result)
	}
}

func printHelp() {
	fmt.Println("Super strange calc 1.0")
	fmt.Println("commands: exchange, geometry, debit")
}

// Программа для конвертации рублей в доллары. Программа запрашивает сумму в рублях и выводит сумму в долларах.
func RubleToDollarExcangeCalculator(rubles float64) float64 {
	return rubles / dollarExcangeRate
}

// Возращает площадь, периметр и гипотенузу прямоугольника.
func GeometryCalc(sideA float64, sideB float64) (float64, float64, float64) {
	sideC := math.Sqrt(math.Pow(sideA, 2) + math.Pow(sideB, 2))
	area := sideA * sideB / 2
	perimeter := sideA + sideB + sideC
	return area, perimeter, sideC
}

// Пользователь вводит сумму вклада в банк и годовой процент. Найти сумму вклада через 5 лет.
func DebitCalc(debit float64, profitPercent float64) float64 {
	return debit + debit*profitPercent/100*5
}
