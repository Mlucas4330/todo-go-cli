package util

import "fmt"

func FormatCurrency(cents int64) string {
	real := float64(cents) / 100
	return fmt.Sprintf("R$ %.2f", real)
}
