package main

import (
	big "math/big"
)

func Fibo(fn_1 big.Int, fn_2 big.Int, result *[]big.Int) {
	var fn big.Int
	fn.Add(&fn_1, &fn_2)
	*result = append(*result, fn_1)
	if len(*result) < cap(*result) {
		Fibo(fn, fn_1, result)
	}
}
