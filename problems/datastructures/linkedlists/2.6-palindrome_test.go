package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterativePalindrom(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{"malayalam", "MalayalaM", true},
		{"empty", "", true},
		{"paap", "paap", true},
		{"p", "p", true},
		{"thing", "thing", false},
		{"Malayalam", "Malayalam", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IterativePalindrom(tt.str); got != tt.want {
				t.Errorf("IterativePalindrom(%s) = %v, want %v", tt.str, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	testcases := []struct {
		name   string
		input  []int
		output bool
	}{
		{"Empty", []int{}, true},
		{"Single Element", []int{1}, true},
		{"Single Element", []int{200}, true},
		{"N ele", []int{200, 3}, false},
		{"N ele", []int{200, 3, 4, 1, 2}, false},
		{"N ele", []int{1, 2, 1}, true},
		{"N ele", []int{1, 2, 2, 1}, true},
		{"N ele", []int{1, 3, 2, 1}, false},
	}
	testFuncs := []func([]int) bool{isPalindrome}
	for _, fn := range testFuncs {
		for _, tc := range testcases {
			t.Run(tc.name, func(t *testing.T) {
				assert.Equal(t, tc.output, fn(tc.input))
			})
		}
	}
}
