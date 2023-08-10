package main

import (
	"testing"
)

func Test_fibonacci(t *testing.T) {
	wants := []int{0, 1, 1, 2, 3, 5}
	f := fibonacci()
	for _, want := range wants {
		if got := f(); got != want {
			t.Errorf("f() = %v, want %v", got, want)
		}
	}
}
