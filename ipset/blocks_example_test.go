// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
)

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
