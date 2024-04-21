package network_test

import (
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
)

func ExampleIterator() {
	first := ip.V4().MustFromBytes(192, 168, 0, 1)
	last := ip.V4().MustFromBytes(192, 168, 0, 254)
	assignable := network.NewRange(first, last)

	// iterator of addresses
	next := assignable.Addresses()
	for address, exists := next(); exists; address, exists = next() {
		println(address.String())
	}
}
