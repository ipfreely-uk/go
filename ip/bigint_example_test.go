package ip_test

import (
	"math/big"
	"testing"

	"github.com/dustin/go-humanize"
	"github.com/ipfreely-uk/go/ip"
)

func TestExampleToBigInt(t *testing.T) {
	ExampleToBigInt()
}

func ExampleToBigInt() {
	printNumberOfAddresses(ip.MinAddress(ip.V4()), ip.MaxAddress(ip.V4()))
	printNumberOfAddresses(ip.MinAddress(ip.V6()), ip.MaxAddress(ip.V6()))
}

func printNumberOfAddresses[A ip.Int[A]](first, last A) {
	diff := last.Subtract(first)

	n := ip.ToBigInt(diff)
	rangeSize := n.Add(n, big.NewInt(1))

	println(first.String(), "to", last.String(), "=", humanize.BigComma(rangeSize), "addresses")
}
