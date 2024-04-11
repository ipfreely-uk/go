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

		ok, actual := nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		ok, actual = nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst1, actual.First())
		assert.Equal(t, expectedLast1, actual.Last())

		ok, _ = nextBlock()
		assert.False(t, ok)
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

		ok, actual := nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		ok, actual = nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst1, actual.First())
		assert.Equal(t, expectedLast1, actual.Last())

		ok, _ = nextBlock()
		assert.False(t, ok)
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

		ok, actual := nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		ok, actual = nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst1, actual.First())
		assert.Equal(t, expectedLast1, actual.Last())

		ok, _ = nextBlock()
		assert.False(t, ok)
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

		ok, actual := nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		ok, actual = nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst1, actual.First())
		assert.Equal(t, expectedLast1, actual.Last())

		ok, _ = nextBlock()
		assert.False(t, ok)
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

		ok, actual := nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		ok, actual = nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst1, actual.First())
		assert.Equal(t, expectedLast1, actual.Last())

		ok, _ = nextBlock()
		assert.False(t, ok)
	}
	{
		family := ip.V4()
		expectedFirst0, _ := ip.Parse(family, "127.0.0.1")
		mask := subnet.Mask(family, 32)
		inverse := mask.Not()
		expectedLast0 := inverse.Or(expectedFirst0)
		input := network.NewRange(expectedFirst0, expectedLast0)

		nextBlock := network.Blocks(input)

		ok, actual := nextBlock()
		assert.True(t, ok)
		assert.Equal(t, expectedFirst0, actual.First())
		assert.Equal(t, expectedLast0, actual.Last())

		ok, _ = nextBlock()
		assert.False(t, ok)
	}
}
