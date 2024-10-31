package ipset_test

import (
	"testing"

	"github.com/dustin/go-humanize"
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
)

func TestExampleParseCIDRNotation(t *testing.T) {
	ExampleParseCIDRNotation()
}

func ExampleParseCIDRNotation() {
	address, mask, _ := ipset.ParseCIDRNotation(ip.V6(), "2001:db8::/32")
	reservedForDocumentation := ipset.NewBlock(address, mask)
	printRangeDetails(reservedForDocumentation)
}

func TestExampleParseUnknownCIDRNotation(t *testing.T) {
	ExampleParseUnknownCIDRNotation()
}

func ExampleParseUnknownCIDRNotation() {
	reservedForDocumentation := []string{
		"192.0.2.0/24",
		"198.51.100.0/24",
		"203.0.113.0/24",
		"2001:db8::/32",
	}
	for _, notation := range reservedForDocumentation {
		address, mask, err := ipset.ParseUnknownCIDRNotation(notation)
		if err != nil {
			panic(err)
		}
		switch a := address.(type) {
		case ip.Addr4:
			printRangeDetails(ipset.NewBlock(a, mask))
		case ip.Addr6:
			printRangeDetails(ipset.NewBlock(a, mask))
		}
	}
}

func printRangeDetails[A ip.Number[A]](addresses ipset.Interval[A]) {
	println("Start:", addresses.First().String())
	println("End:", addresses.Last().String())
	println("Addresses:", humanize.BigComma(addresses.Size()))
	println()
}
