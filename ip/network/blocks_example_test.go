package network_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
)

func TestExampleBlocks(t *testing.T) {
	ExampleBlocks()
}

func ExampleBlocks() {
	first := ip.V4().MustFromBytes(192, 0, 2, 101)
	last := ip.V4().MustFromBytes(192, 0, 2, 240)
	freeAddresses := network.NewRange(first, last)

	printCidrBlocksIn(freeAddresses)
}

func printCidrBlocksIn[A ip.Number[A]](addressRange network.AddressRange[A]) {
	next := network.Blocks(addressRange)
	for block, exists := next(); exists; block, exists = next() {
		println(block.String())
	}
}
