package ip_test

import (
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

func ExampleToBigInt() {
	family := ip.V6()
	first := ip.MinAddress(family)
	last := ip.MaxAddress(family)
	diff := last.Subtract(first)

	n := ip.ToBigInt(diff)
	rangeSize := n.Add(n, big.NewInt(1))

	println("Number of addresses = ", rangeSize.String())
}
