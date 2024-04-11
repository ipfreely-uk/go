package cidr_test

import (
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/ipfreely-uk/go/ip/network/cidr"
)

func ExampleNotation() {
	address, _ := ip.V4().FromBytes(192, 168, 0, 1)
	single := network.NewBlock(address, address.Family().Width())
	notation := cidr.Notation(single)
	println(notation)
}

func ExampleParse() {
	reservedForDocumentation, _ := cidr.Parse(ip.V6(), "2001:db8::/32")
	printRangeDetails(reservedForDocumentation)
}

func ExampleParseUnknown() {
	reservedForDocumentation := []string{
		"192.0.2.0/24",
		"198.51.100.0/24",
		"203.0.113.0/24",
		"2001:db8::/32",
	}
	for _, notation := range reservedForDocumentation {
		block, err := cidr.ParseUnknown(notation)
		if err != nil {
			panic(err)
		}
		switch addresses := block.(type) {
		case network.Block[ip.A4]:
			printRangeDetails(addresses)
		case network.Block[ip.A6]:
			printRangeDetails(addresses)
		}
	}
}

func printRangeDetails[A ip.Address[A]](addresses network.AddressRange[A]) {
	println("Start:", addresses.First().String())
	println("End:", addresses.Last().String())
	println("Addresses:", addresses.Size().String())
	println()
}
