package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		input string
		want  map[string]int
	}{
		{"I am learning Go!", map[string]int{
			"I": 1, "am": 1, "learning": 1, "Go!": 1,
		}},
		{"The quick brown fox jumped over the lazy dog.", map[string]int{
			"The": 1, "quick": 1, "brown": 1, "fox": 1, "jumped": 1,
			"over": 1, "the": 1, "lazy": 1, "dog.": 1,
		}},
		{"I ate a donut. Then I ate another donut.", map[string]int{
			"I": 2, "ate": 2, "a": 1, "donut.": 2, "Then": 1, "another": 1,
		}},
		{"A man a plan a canal panama.", map[string]int{
			"A": 1, "man": 1, "a": 2, "plan": 1, "canal": 1, "panama.": 1,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := WordCount(tt.input)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("WordCount()  mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
