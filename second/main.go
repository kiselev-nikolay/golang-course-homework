package main

import (
	"fmt"
	"strconv"
)

func IsOdd(number int) bool {
	return number%2 == 0
}

func IsTriple(number int) bool {
	return number%3 == 0
}

func main() {
	fmt.Print("Any integer number: ")
	var user_input string
	fmt.Scanln(&user_input)
	number, err := strconv.Atoi(user_input)
	if err != nil {
		fmt.Printf("'%v' is not int.\n", user_input)
		fmt.Println(err)
		return
	}
	if IsOdd(number) {
		fmt.Printf("'%v' is odd.\n", number)
	} else {
		fmt.Printf("'%v' is even.\n", number)
	}
	if IsTriple(number) {
		fmt.Printf("'%v' is divisible by three.\n", number)
	} else {
		fmt.Printf("'%v' is not divisible by three.\n", number)
	}
}
