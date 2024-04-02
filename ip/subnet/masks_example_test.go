package subnet_test

import (
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/subnet"
)

func Example_subnet_Mask() {
	network, _ := ip.V4().FromBytes(192, 168, 0, 0)
	mask := subnet.Mask(ip.V4(), 24)

	println("First: %s", network.String())
	println("Last: %s", mask.Not().Or(network).String())
	println("Mask: %s", mask.String())
}
