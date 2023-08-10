package main

import "testing"

func TestMyReader_Read(t *testing.T) {
	r := MyReader{}

	b := make([]byte, 1024, 2048)
	i, o := 0, 0
	for ; i < 1<<20 && o < 1<<20; i++ {
		n, err := r.Read(b)
		for i, v := range b[:n] {
			if v != 'A' {
				t.Fatalf("got byte %x at offset %v, want 'A'\n", v, o+i)
			}
		}
		o += n
		if err != nil {
			t.Fatalf("read error: %v\n", err)
		}
	}
	if o == 0 {
		t.Fatalf("read zero bytes after %d Read calls\n", i)
	}
}
