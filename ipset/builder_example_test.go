package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipmask"
	"github.com/ipfreely-uk/go/ipset"
)

func TestExampleBuilder(t *testing.T) {
	ExampleBuilder()
}

func ExampleBuilder() {
	addr, bits, err := ipmask.ParseCIDRNotation(ip.V4(), "10.0.0.0/15")
	if err != nil {
		panic(err.Error())
	}
	network := ipset.NewBlock(addr, bits)

	lucky := removeUnlucky(network)

	println(lucky.String())
}

var thirteen = ip.V4().FromInt(13)
var maskComplement = ipmask.For(ip.V4(), 24).Not()

func removeUnlucky(addresses ipset.Discrete[ip.Addr4]) (lucky ipset.Discrete[ip.Addr4]) {
	bldr := ipset.Builder[ip.Addr4]{}
	for a := range addresses.Addresses() {
		if !unlucky(a) {
			bldr.Union(ipset.NewSingle(a))
		}
	}
	return bldr.Build()
}

func unlucky(a ip.Addr4) bool {
	lastDigits := a.And(maskComplement)
	return ip.Eq(lastDigits, thirteen)
}
