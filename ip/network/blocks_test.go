package network_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/ipfreely-uk/go/ip/subnet"
	"github.com/stretchr/testify/assert"
)

func TestBlocks(t *testing.T) {
	{
		family := ip.V6()
		expectedFirst0, _ := ip.Parse(family, "fe80::")
		mask := subnet.Mask(family, 64)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := network.NewRange(expectedFirst0, expectedLast1)

		nextBlock := network.Blocks(input)

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
		mask := subnet.Mask(family, 8)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := network.NewRange(expectedFirst0, expectedLast1)

		nextBlock := network.Blocks(input)

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
		mask := subnet.Mask(family, 128)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := network.NewRange(expectedFirst0, expectedLast1)

		nextBlock := network.Blocks(input)

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
		mask := subnet.Mask(family, 8)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := network.NewRange(expectedFirst0, expectedLast1)

		nextBlock := network.Blocks(input)

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
		mask := subnet.Mask(family, 32)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		expectedFirst1 := ip.Next(expectedLast0)
		expectedLast1 := expectedFirst1
		input := network.NewRange(expectedFirst0, expectedLast1)

		nextBlock := network.Blocks(input)

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
		mask := subnet.Mask(family, 32)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		input := network.NewRange(expectedFirst0, expectedLast0)

		nextBlock := network.Blocks(input)

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
		input := network.NewRange(first, last)

		next := network.Blocks(input)

		for block, exists := next(); exists; block, exists = next() {
			assert.True(t, input.Contains(block.First()))
			assert.True(t, input.Contains(block.Last()))
		}
	}
}
