package cardat

import (
	"fmt"
)

func cardAt(n int) string {
	if n < 0 || n > 51 {
		return "Out of range"
	}
	suits := []byte{'C', 'D', 'H', 'S'}
	ranks := []byte{'2', '3', '4', '5', '6', '7', '8', '9', 'O', 'J', 'Q', 'K', 'A'}
	return fmt.Sprintf("%c%c", ranks[n%13], suits[n/13])
}
