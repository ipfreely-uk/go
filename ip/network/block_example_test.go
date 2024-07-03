package network_test

import (
	"crypto/rand"
	"testing"

	humanize "github.com/dustin/go-humanize"
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
)

func TestExampleNewBlock(t *testing.T) {
	ExampleNewBlock()
}

func ExampleNewBlock() {
	netAddress := ip.MustParse(ip.V6(), "2001:db8::")

	block := network.NewBlock(netAddress, 32)

	println("Block", block.String())
	println("First", block.First().String())
	println("Last", block.Last().String())
	println("Size", humanize.BigComma(block.Size()))
}

func ExampleBlock() {
	netAddress := ip.MustParse(ip.V6(), "2001:db8:cafe::")
	block := network.NewBlock(netAddress, 56)

	randomAddr := randomAddressFrom(block)

	println("Random address from", block.String(), "=", randomAddr.String())
}

func randomAddressFrom[A ip.Number[A]](netBlock network.Block[A]) (address A) {
	netAddr := netBlock.First()
	family := netAddr.Family()
	inverseMask := netBlock.Mask().Not()

	return random(family).And(inverseMask).Or(netAddr)
}

func random[A ip.Number[A]](f ip.Family[A]) (address A) {
	slice := make([]byte, f.Width()/8)
	_, _ = rand.Read(slice)
	return f.MustFromBytes(slice...)
}
