package main

import (
	"fmt"
	"strconv"
)

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
	if number%2 == 0 {
		fmt.Printf("'%v' is odd.\n", number)
	} else {
		fmt.Printf("'%v' is even.\n", number)
	}
}
