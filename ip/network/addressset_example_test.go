package network_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
)

func TestExampleNewSet(t *testing.T) {
	ExampleNewSet()
}

func ExampleNewSet() {
	family := ip.V4()
	r0 := exampleRange(family, "192.0.2.0", "192.0.2.100")
	r1 := exampleRange(family, "192.0.2.101", "192.0.2.111")
	r2 := exampleRange(family, "192.0.2.200", "192.0.2.255")
	r3 := exampleRange(family, "203.0.113.0", "203.0.113.255")
	r4 := exampleRange(family, "192.0.2.0", "192.0.2.100")

	addresses := network.NewSet(r0, r1, r2, r3, r4)

	println("Rationalized ranges:")
	next := addresses.Ranges()
	for aRange, exists := next(); exists; aRange, exists = next() {
		println(aRange.String())
	}
}

func exampleRange[A ip.Number[A]](family ip.Family[A], first, last string) network.AddressRange[A] {
	a0 := ip.MustParse(family, first)
	a1 := ip.MustParse(family, last)
	return network.NewRange(a0, a1)
}
