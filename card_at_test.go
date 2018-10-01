package cardat

import (
	"fmt"
	"testing"
)

func TestCardAt(t *testing.T) {
	testCases := []struct {
		n        int
		expected string
	}{
		{
			n:        -1,
			expected: "Out of range",
		},
		{
			n:        0,
			expected: "2C",
		},
		{
			n:        1,
			expected: "3C",
		},
		{
			n:        34,
			expected: "OH",
		},
		{
			n:        35,
			expected: "JH",
		},
		{
			n:        52,
			expected: "Out of range",
		},
	}
	for _, tC := range testCases {
		desc := fmt.Sprintf("cardAt should return %s when n = %d", tC.expected, tC.n)
		t.Run(desc, func(t *testing.T) {
			actual := cardAt(tC.n)
			if actual != tC.expected {
				t.Errorf("expected %s but got %s", tC.expected, actual)
			}
		})
	}
}
