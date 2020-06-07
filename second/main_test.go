package main

import (
	"testing"
)

func TestIsOdd(t *testing.T) {
	got := IsOdd(3)
	if got {
		t.Errorf("IsOdd(3) = %v; want false", got)
	}
	got = IsOdd(12)
	if !got {
		t.Errorf("IsOdd(12) = %v; want true", got)
	}
}

func TestIsTriple(t *testing.T) {
	got := IsTriple(31)
	if got {
		t.Errorf("IsTriple(31) = %v; want false", got)
	}
	got = IsTriple(33)
	if !got {
		t.Errorf("IsTriple(33) = %v; want true", got)
	}
}
