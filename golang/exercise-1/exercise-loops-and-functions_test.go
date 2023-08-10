package main

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSqrt(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "Sqrt(2)", args: args{x: 2}, want: math.Sqrt(2)},
		{name: "Sqrt(3)", args: args{x: 3}, want: math.Sqrt(3)},
	}
	opt := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return delta/mean < 0.00001
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sqrt(tt.args.x); !cmp.Equal(tt.want, got, opt) {
				t.Errorf("Sqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
