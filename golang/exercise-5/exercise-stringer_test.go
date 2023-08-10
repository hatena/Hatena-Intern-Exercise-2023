package main

import (
	"fmt"
	"testing"
)

func TestIPAddr_String(t *testing.T) {
	tests := []struct {
		name string
		addr IPAddr
		want string
	}{
		{"loopback", IPAddr{127, 0, 0, 1}, "127.0.0.1"},
		{"googleDNS", IPAddr{8, 8, 8, 8}, "8.8.8.8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fmt.Sprint(tt.addr); got != tt.want {
				t.Errorf("IPAddr.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
