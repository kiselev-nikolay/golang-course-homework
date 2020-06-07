package main

import (
	"fmt"
	big "math/big"
	"testing"
)

func TestFibo(t *testing.T) {
	var fn_1 big.Int
	fn_1.SetInt64(0)
	var fn_2 big.Int
	fn_2.SetInt64(1)
	got := make([]big.Int, 0, 101)
	Fibo(fn_1, fn_2, &got)
	fmt.Println(got)
	t.Error(got[100].String() == "354224848179261915075")
}
