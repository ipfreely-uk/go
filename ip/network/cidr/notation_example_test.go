package cidr_test

import (
	"github.com/dustin/go-humanize"
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/ipfreely-uk/go/ip/network/cidr"
)

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
		case network.Block[ip.Addr4]:
			printRangeDetails(addresses)
		case network.Block[ip.Addr6]:
			printRangeDetails(addresses)
		}
	}
}

func printRangeDetails[A ip.Address[A]](addresses network.AddressRange[A]) {
	println("Start:", addresses.First().String())
	println("End:", addresses.Last().String())
	println("Addresses:", humanize.BigComma(addresses.Size()))
	println()
}
