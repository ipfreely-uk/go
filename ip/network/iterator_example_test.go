package network_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
)

func TestExampleIterator(t *testing.T) {
	ExampleIterator()
}

func ExampleIterator() {
	netaddr, mask, _ := network.ParseCIDRNotation(ip.V4(), "192.0.2.128/28")
	subnet := network.NewBlock(netaddr, mask)
	firstAssigneable := ip.Next(subnet.First())
	lastAssigneable := ip.Prev(subnet.Last())
	assignable := network.NewRange(firstAssigneable, lastAssigneable)

	// iterator of addresses
	next := assignable.Addresses()
	for address, exists := next(); exists; address, exists = next() {
		println(address.String())
	}
}
