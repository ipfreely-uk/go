// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipmask"
	"github.com/ipfreely-uk/go/ipset"
	"github.com/ipfreely-uk/go/txt"
)

func TestExampleSubnets(t *testing.T) {
	ExampleSubnets()
}

func maskRequiredFor[A ip.Int[A]](f ip.Family[A], allocateableAddresses *big.Int) (bits int) {
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
		sizeForMask := ipmask.Size(f, m)
		if sizeForMask.Cmp(min) >= 0 {
			return m
		}
	}
	formatted := txt.CommaDelim(allocateableAddresses)
	msg := fmt.Sprintf("%s is larger than family %s", formatted, f.String())
	panic(msg)
}

func ExampleSubnets() {
	oneHundredAddresses := big.NewInt(100)
	mask := maskRequiredFor(ip.V4(), oneHundredAddresses)

	netAddr, bits, _ := ipset.ParseCIDRNotation(ip.V4(), "203.0.113.0/24")
	block := ipset.NewBlock(netAddr, bits)

	println(fmt.Sprintf("Dividing %s into blocks of at least %s addresses", block.CidrNotation(), oneHundredAddresses.String()))
	for sub := range ipset.Subnets(block, mask) {
		println(sub.String())
	}
}

func TestExampleBlocks(t *testing.T) {
	ExampleBlocks()
}

func ExampleBlocks() {
	first := ip.V4().MustFromBytes(192, 0, 2, 101)
	last := ip.V4().MustFromBytes(192, 0, 2, 240)
	freeAddresses := ipset.NewInterval(first, last)

	printCidrBlocksIn(freeAddresses)
}

func printCidrBlocksIn[A ip.Int[A]](addressRange ipset.Interval[A]) {
	for block := range ipset.Blocks(addressRange) {
		println(block.String())
	}
}
