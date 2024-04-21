package network_test

import (
	"crypto/rand"

	humanize "github.com/dustin/go-humanize"
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
)

func ExampleNewBlock() {
	netAddress, _ := ip.Parse(ip.V6(), "2001:db8::")

	block := network.NewBlock(netAddress, 32)

	println("Block", block.String())
	println("First", block.First().String())
	println("Last", block.Last().String())
	println("Size", humanize.BigComma(block.Size()))
}

func ExampleBlock_Mask() {
	netAddress, _ := ip.Parse(ip.V6(), "2001:db8:cafe::")
	block := network.NewBlock(netAddress, 56)

	randomAddr := randomAddressFrom(block)

	println("Random address from", block.String(), "=", randomAddr.String())
}

func randomAddressFrom[A ip.Address[A]](netBlock network.Block[A]) A {
	netAddr := netBlock.First()
	family := netAddr.Family()
	r := randomAddress(family)

	mask := netBlock.Mask()

	return r.Xor(mask).Or(netAddr)
}

func randomAddress[A ip.Address[A]](f ip.Family[A]) A {
	slice := make([]byte, f.Width()/8)
	_, _ = rand.Read(slice)
	return f.MustFromBytes(slice...)
}
