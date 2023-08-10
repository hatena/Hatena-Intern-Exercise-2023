package main

import (
	"errors"
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
		err  error
	}{
		{"positive", args{2.0}, math.Sqrt(2.0), nil},
		{"negative", args{-2.0}, math.Sqrt(-2.0), errors.New("cannot Sqrt negative number: -2")},
	}
	opt := cmp.Comparer(func(x, y float64) bool {
		delta := math.Abs(x - y)
		mean := math.Abs(x+y) / 2.0
		return delta/mean < 0.00001
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Sqrt(tt.args.x)
			if tt.err == nil {
				if err != nil {
					t.Fatalf("Sqrt(%v) want no error, but got error %v", tt.args.x, err)
				} else if !cmp.Equal(tt.want, got, opt) {
					t.Errorf("Sqrt() = %v, want %v", got, tt.want)
				}
			} else if err == nil || err.Error() != tt.err.Error() {
				t.Fatalf("Sqrt(%v) want return error \"%v\", but got \"%v\"", tt.args.x, tt.err, err)

			}
		})
	}
}
