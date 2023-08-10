package main

import (
	"bytes"
	"image/png"
	"testing"
)

func TestImage(t *testing.T) {
	m := Image{}
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		t.Error(err)
	}
}
