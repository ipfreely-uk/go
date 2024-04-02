package network_test

import (
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
)

func Example_network_Iterator() {
	first, _ := ip.V4().FromBytes(192, 168, 0, 1)
	last, _ := ip.V4().FromBytes(192, 168, 0, 254)
	assignable := network.NewRange(first, last)
	addrs := assignable.Addresses()
	for ok, address := addrs(); ok; ok, address = addrs() {
		println(address.String())
	}
}
