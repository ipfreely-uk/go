package network_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	f := ip.V6()
	min := ip.MinAddress(f)
	one := f.FromInt(1)
	two := f.FromInt(2)
	ten := f.FromInt(10)
	hundred := f.FromInt(100)

	one_one := network.NewBlock(one, 128)
	two_two := network.NewBlock(two, 128)
	one_ten := network.NewRange(one, ten)
	two_hundred := network.NewRange(two, hundred)
	internet := network.NewBlock(min, 0)

	assert.True(t, network.Intersect(one_one, one_one))
	assert.True(t, network.Intersect(one_one, one_ten))
	assert.True(t, network.Intersect(one_ten, one_one))
	assert.True(t, network.Intersect(one_ten, two_hundred))
	assert.True(t, network.Intersect(internet, two_hundred))
	assert.True(t, network.Intersect(two_hundred, two_hundred))
	assert.False(t, network.Intersect(one_one, two_two))
	assert.False(t, network.Intersect(one_one, two_hundred))
}

func TestAdjacent(t *testing.T) {
	f := ip.V6()
	min := ip.MinAddress(f)
	one := f.FromInt(1)
	two := f.FromInt(2)
	ten := f.FromInt(10)
	hundred := f.FromInt(100)
	hundredOne := f.FromInt(101)

	one_one := network.NewBlock(one, 128)
	two_two := network.NewBlock(two, 128)
	one_ten := network.NewRange(one, ten)
	two_hundred := network.NewRange(two, hundred)
	hundredOne_hundredOne := network.NewBlock(hundredOne, 128)
	internet := network.NewBlock(min, 0)

	assert.False(t, network.Adjacent(one_one, one_one))
	assert.False(t, network.Adjacent(one_one, one_ten))
	assert.False(t, network.Adjacent(one_ten, one_one))
	assert.False(t, network.Adjacent(one_ten, two_hundred))
	assert.False(t, network.Adjacent(internet, two_hundred))
	assert.False(t, network.Adjacent(two_hundred, two_hundred))
	assert.True(t, network.Adjacent(one_one, two_two))
	assert.True(t, network.Adjacent(one_one, two_hundred))
	assert.True(t, network.Adjacent(two_hundred, one_one))
	assert.True(t, network.Adjacent(hundredOne_hundredOne, two_hundred))
	assert.True(t, network.Adjacent(two_hundred, hundredOne_hundredOne))
}

func TestContiguous(t *testing.T) {
	f := ip.V6()
	min := ip.MinAddress(f)
	one := f.FromInt(1)
	two := f.FromInt(2)
	ten := f.FromInt(10)
	hundred := f.FromInt(100)
	hundredOne := f.FromInt(101)

	one_one := network.NewBlock(one, 128)
	two_two := network.NewBlock(two, 128)
	one_ten := network.NewRange(one, ten)
	two_hundred := network.NewRange(two, hundred)
	hundredOne_hundredOne := network.NewBlock(hundredOne, 128)
	internet := network.NewBlock(min, 0)

	assert.True(t, network.Contiguous(one_one, one_one))
	assert.True(t, network.Contiguous(one_one, one_ten))
	assert.True(t, network.Contiguous(one_ten, one_one))
	assert.True(t, network.Contiguous(one_ten, two_hundred))
	assert.True(t, network.Contiguous(internet, two_hundred))
	assert.True(t, network.Contiguous(two_hundred, two_hundred))
	assert.True(t, network.Contiguous(one_one, two_two))
	assert.True(t, network.Contiguous(one_one, two_hundred))
	assert.True(t, network.Contiguous(two_hundred, one_one))
	assert.True(t, network.Contiguous(hundredOne_hundredOne, two_hundred))
	assert.True(t, network.Contiguous(two_hundred, hundredOne_hundredOne))
	assert.False(t, network.Contiguous(two_two, hundredOne_hundredOne))
	assert.False(t, network.Contiguous(hundredOne_hundredOne, two_two))
}

func TestJoin(t *testing.T) {
	f := ip.V6()
	min := ip.MinAddress(f)
	one := f.FromInt(1)
	hundredOne := f.FromInt(101)

	one_one := network.NewBlock(one, 128)
	hundredOne_hundredOne := network.NewBlock(hundredOne, 128)
	internet := network.NewBlock(min, 0)

	{
		actual := network.Join(one_one, hundredOne_hundredOne)
		assert.Equal(t, one, actual.First())
		assert.Equal(t, hundredOne, actual.Last())
	}
	{
		actual := network.Join(hundredOne_hundredOne, one_one)
		assert.Equal(t, one, actual.First())
		assert.Equal(t, hundredOne, actual.Last())
	}
	{
		actual := network.Join(internet, one_one)
		assert.Same(t, internet, actual)
	}
	{
		actual := network.Join(one_one, internet)
		assert.Same(t, internet, actual)
	}
}
