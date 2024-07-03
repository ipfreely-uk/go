package network_test

import (
	"crypto/rand"
	"testing"

	humanize "github.com/dustin/go-humanize"
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/compare"
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

func TestExampleBlock(t *testing.T) {
	ExampleBlock()
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

func TestExampleBlock_second(t *testing.T) {
	ExampleBlock_second()
}

func ExampleBlock_second() {
	netAddress := ip.MustParse(ip.V6(), "2001:db8:cafe::")
	block := network.NewBlock(netAddress, 56)

	next := split(block, 60)
	for subnet, exists := next(); exists; subnet, exists = next() {
		println(subnet.String())
	}
}

// Split subnet into smaller subnets
func split[A ip.Number[A]](b network.Block[A], newMaskSize int) network.Iterator[network.Block[A]] {
	if newMaskSize < b.MaskSize() {
		panic("invalid split size")
	}
	current := network.NewBlock(b.First(), newMaskSize)
	one := current.First().Family().FromInt(1)
	increment := current.Last().Subtract(current.First()).Add(one)
	exhausted := false

	return func() (element network.Block[A], exists bool) {
		if exhausted {
			return nil, false
		}
		result := current
		exhausted = compare.Eq(current.Last(), b.Last())
		if !exhausted {
			first := current.First().Add(increment)
			current = network.NewBlock(first, newMaskSize)
		}
		return result, true
	}
}
