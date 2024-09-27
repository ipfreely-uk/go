package ipset_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/ipset"
)

func TestExampleSubnets(t *testing.T) {
	ExampleSubnets()
}

func maskRequiredFor[A ip.Number[A]](f ip.Family[A], allocateableAddresses *big.Int) (bits int) {
	var min *big.Int
	if f.Version() == ip.V4().Version() {
		// IPv4 subnets reserve two addresses for the network
		two := big.NewInt(2)
		min = two.Add(two, allocateableAddresses)
	} else {
		min = allocateableAddresses
	}
	width := f.Width()
	for m := width; m >= 0; m-- {
		sizeForMask := ip.SubnetAddressCount(f, m)
		if sizeForMask.Cmp(min) >= 0 {
			return m
		}
	}
	msg := fmt.Sprintf("%s is larger than family %s", allocateableAddresses.String(), f.String())
	panic(msg)
}

func ExampleSubnets() {
	oneHundredAddresses := big.NewInt(100)
	mask := maskRequiredFor(ip.V4(), oneHundredAddresses)

	netAddr, bits, _ := ipset.ParseCIDRNotation(ip.V4(), "203.0.113.0/24")
	block := ipset.NewBlock(netAddr, bits)

	println(fmt.Sprintf("Dividing %s into blocks of at least %s addresses", block.CidrNotation(), oneHundredAddresses.String()))
	nextSub := ipset.Subnets(block, mask)
	for sub, exists := nextSub(); exists; sub, exists = nextSub() {
		println(sub.String())
	}
}
