// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset_test

import (
	"iter"
	"math/rand"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
	"github.com/stretchr/testify/assert"
)

func TestSubnets(t *testing.T) {
	netAddr := ip.MustParse(ip.V4(), "192.168.0.0")
	{
		b := ipset.NewBlock(netAddr, 31)
		next, stop := iter.Pull(ipset.Subnets(b, 32))
		defer stop()

		b, exists := next()
		assert.True(t, exists)
		assert.Equal(t, "192.168.0.0/32", b.String())

		b, exists = next()
		assert.True(t, exists)
		assert.Equal(t, "192.168.0.1/32", b.String())

		_, exists = next()
		assert.False(t, exists)
	}
	{
		b := ipset.NewBlock(netAddr, 31)
		assert.Panics(t, func() {
			ipset.Subnets(b, 12)
		})
		assert.Panics(t, func() {
			ipset.Subnets(b, 33)
		})
	}
}

func TestBlocks(t *testing.T) {
	{
		family := ip.V6()
		expectedFirst0, _ := ip.Parse(family, "fe80::")
		mask := ip.SubnetMask(family, 64)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := ipset.NewInterval(expectedFirst0, expectedLast1)

		nextBlock, stop := iter.Pull(ipset.Blocks(input))
		defer stop()

		actual, exists := nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		actual, exists = nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst1, actual.First())
		assert.Equal(t, expectedLast1, actual.Last())

		_, exists = nextBlock()
		assert.False(t, exists)
	}
	{
		family := ip.V6()
		expectedFirst0, _ := ip.Parse(family, "f000::")
		mask := ip.SubnetMask(family, 8)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := ipset.NewInterval(expectedFirst0, expectedLast1)

		nextBlock, stop := iter.Pull(ipset.Blocks(input))
		defer stop()

		actual, exists := nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		actual, exists = nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst1, actual.First())
		assert.Equal(t, expectedLast1, actual.Last())

		_, exists = nextBlock()
		assert.False(t, exists)
	}
	{
		family := ip.V6()
		expectedFirst0, _ := ip.Parse(family, "::1")
		mask := ip.SubnetMask(family, 128)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := ipset.NewInterval(expectedFirst0, expectedLast1)

		nextBlock, stop := iter.Pull(ipset.Blocks(input))
		defer stop()

		actual, exists := nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		actual, exists = nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst1, actual.First())
		assert.Equal(t, expectedLast1, actual.Last())

		_, exists = nextBlock()
		assert.False(t, exists)
	}
	{
		family := ip.V4()
		expectedFirst0, _ := ip.Parse(family, "10.0.0.0")
		mask := ip.SubnetMask(family, 8)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := ipset.NewInterval(expectedFirst0, expectedLast1)

		nextBlock, stop := iter.Pull(ipset.Blocks(input))
		defer stop()

		actual, exists := nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		actual, exists = nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst1, actual.First())
		assert.Equal(t, expectedLast1, actual.Last())

		_, exists = nextBlock()
		assert.False(t, exists)
	}
	{
		family := ip.V4()
		expectedFirst0, _ := ip.Parse(family, "127.0.0.1")
		mask := ip.SubnetMask(family, 32)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := ipset.NewInterval(expectedFirst0, expectedLast1)

		nextBlock, stop := iter.Pull(ipset.Blocks(input))
		defer stop()

		actual, exists := nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		actual, exists = nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst1, actual.First())
		assert.Equal(t, expectedLast1, actual.Last())

		_, exists = nextBlock()
		assert.False(t, exists)
	}
	{
		family := ip.V4()
		expectedFirst0, _ := ip.Parse(family, "127.0.0.1")
		mask := ip.SubnetMask(family, 32)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		input := ipset.NewInterval(expectedFirst0, expectedLast0)

		nextBlock, stop := iter.Pull(ipset.Blocks(input))
		defer stop()

		actual, exists := nextBlock()
		assert.True(t, exists)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		_, exists = nextBlock()
		assert.False(t, exists)
	}
	{
		first := ip.V6().FromInt(999)
		last := ip.V6().FromInt(0).Not()
		input := ipset.NewInterval(first, last)

		for block := range ipset.Blocks(input) {
			assert.True(t, input.Contains(block.First()))
			assert.True(t, input.Contains(block.Last()))
		}
	}
	{
		family := ip.V6()
		expectedFirst0, _ := ip.Parse(family, "fe80::")
		mask := ip.SubnetMask(family, 64)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := ipset.NewInterval(expectedFirst0, expectedLast1)
		// check algo stops early
		count := 0
		ipset.Blocks[ip.Addr6](input)(func(b ipset.Block[ip.Addr6]) bool {
			count++
			return false
		})

		assert.Equal(t, 1, count)
	}
}

func TestBlockIteration(t *testing.T) {
	makeAndWalkBlocks(t, ip.V4())
	makeAndWalkBlocks(t, ip.V6())
}

func makeAndWalkBlocks[A ip.Int[A]](t *testing.T, family ip.Family[A]) {
	src := rand.New(rand.NewSource(0))
	buf := make([]byte, family.Width()/8)

	prev := family.MustFromBytes(buf...)
	for i := 0; i < 200; i++ {
		_, err := src.Read(buf)
		assert.Nil(t, err)

		next := family.MustFromBytes(buf...)
		walkBlocks(t, prev, next)

		prev = next
	}
}

func walkBlocks[A ip.Int[A]](t *testing.T, a1, a2 A) {
	r := ipset.NewInterval(a1, a2)
	nextBlock, stop := iter.Pull(ipset.Blocks(r))
	defer stop()

	prev, _ := nextBlock()

	for block, exists := nextBlock(); exists; block, exists = nextBlock() {
		assert.True(t, ipset.Adjacent(prev, block))
		assert.False(t, ipset.Intersect(prev, block))

		prev = block
	}
}
