package main

import (
	"fmt"
	"math"
)

type tax struct {
	taxPercent float64
}

func newTax() tax {
	return tax{taxPercent: 0.08}
}

func roundTax(f float64) float64 {
	return math.Floor(f + .5)
}

func (t tax) calculateTaxExcludeAmount(taxIncludePrice int) int {
	price := float64(taxIncludePrice)
	return int(roundTax(price / (1 + t.taxPercent)))
}

func main() {
	tax := newTax()
	fmt.Printf("税込み: %v, 税抜き: %v", 2000, tax.calculateTaxExcludeAmount(2000))
}