package util

import "fmt"

type Invoice struct {
	AmountCents int64
}

func FormatCurrency(cents int64) string {
	real := float64(cents) / 100
	return fmt.Sprintf("R$ %.2f", real)
}
