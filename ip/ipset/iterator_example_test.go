package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/ipset"
)

func TestExampleIterator(t *testing.T) {
	ExampleIterator()
}

func ExampleIterator() {
	netaddr, mask, _ := ipset.ParseCIDRNotation(ip.V4(), "192.0.2.128/28")
	subnet := ipset.NewBlock(netaddr, mask)
	firstAssigneable := ip.Next(subnet.First())
	lastAssigneable := ip.Prev(subnet.Last())
	assignable := ipset.NewInterval(firstAssigneable, lastAssigneable)

	// iterator of addresses
	next := assignable.Addresses()
	for address, exists := next(); exists; address, exists = next() {
		println(address.String())
	}
}

func TestExampleSeq(t *testing.T) {
	ExampleSeq()
}

func ExampleSeq() {
	netaddr, mask, _ := ipset.ParseCIDRNotation(ip.V4(), "192.0.2.128/28")
	subnet := ipset.NewBlock(netaddr, mask)
	firstAssigneable := ip.Next(subnet.First())
	lastAssigneable := ip.Prev(subnet.Last())
	assignable := ipset.NewInterval(firstAssigneable, lastAssigneable)

	for address := range ipset.Seq(assignable.Addresses()) {
		println(address.String())
	}
}
