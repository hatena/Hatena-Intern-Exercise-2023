package main

import (
	"testing"
)

func TestPic(t *testing.T) {
	const (
		dx = 256
		dy = 256
	)
	got := Pic(dx, dy)
	if len(got) != dx {
		t.Fatalf("picture width = %v, want %v", len(got), dx)
	}
	for i := range got {
		if len(got[i]) != dy {
			t.Fatalf("picture height = %v, want %v", len(got[i]), dy)
		}
	}
}
