package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

func Test_rot13Reader_Read(t *testing.T) {
	tests := []struct {
		want  string
		input string
	}{
		{"You cracked the code!", "Lbh penpxrq gur pbqr!"},
		{"The Quick Brown Fox Jumps Over The Lazy Dog", "Gur Dhvpx Oebja Sbk Whzcf Bire Gur Ynml Qbt"},
	}
	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			s := strings.NewReader(tt.input)
			r := rot13Reader{s}
			buf := &bytes.Buffer{}
			_, _ = io.Copy(buf, &r)
			got := buf.String()

			if got != tt.want {
				t.Errorf("rot13Reader.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}
