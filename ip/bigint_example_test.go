package ip_test

import (
	"math/big"

	"github.com/dustin/go-humanize"
	"github.com/ipfreely-uk/go/ip"
)

func ExampleToBigInt() {
	printNumberOfAddresses(ip.MinAddress(ip.V4()), ip.MaxAddress(ip.V4()))
	printNumberOfAddresses(ip.MinAddress(ip.V6()), ip.MaxAddress(ip.V6()))
}

func printNumberOfAddresses[A ip.Address[A]](first, last A) {
	diff := last.Subtract(first)

	n := ip.ToBigInt(diff)
	rangeSize := n.Add(n, big.NewInt(1))

	println(first.String(), "-", last.String(), "=", humanize.BigComma(rangeSize))
}
