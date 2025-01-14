// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	f := ip.V6()
	min := ip.MinAddress(f)
	one := f.FromInt(1)
	two := f.FromInt(2)
	ten := f.FromInt(10)
	hundred := f.FromInt(100)

	one_one := ipset.NewBlock(one, 128)
	two_two := ipset.NewBlock(two, 128)
	one_ten := ipset.NewInterval(one, ten)
	two_hundred := ipset.NewInterval(two, hundred)
	internet := ipset.NewBlock(min, 0)

	assert.True(t, ipset.Intersect(one_one, one_one))
	assert.True(t, ipset.Intersect(one_one, one_ten))
	assert.True(t, ipset.Intersect(one_ten, one_one))
	assert.True(t, ipset.Intersect(one_ten, two_hundred))
	assert.True(t, ipset.Intersect(internet, two_hundred))
	assert.True(t, ipset.Intersect(two_hundred, two_hundred))
	assert.False(t, ipset.Intersect(one_one, two_two))
	assert.False(t, ipset.Intersect(one_one, two_hundred))
}

func TestAdjacent(t *testing.T) {
	f := ip.V6()
	min := ip.MinAddress(f)
	one := f.FromInt(1)
	two := f.FromInt(2)
	ten := f.FromInt(10)
	hundred := f.FromInt(100)
	hundredOne := f.FromInt(101)

	one_one := ipset.NewBlock(one, 128)
	two_two := ipset.NewBlock(two, 128)
	one_ten := ipset.NewInterval(one, ten)
	two_hundred := ipset.NewInterval(two, hundred)
	hundredOne_hundredOne := ipset.NewBlock(hundredOne, 128)
	internet := ipset.NewBlock(min, 0)

	assert.False(t, ipset.Adjacent(one_one, one_one))
	assert.False(t, ipset.Adjacent(one_one, one_ten))
	assert.False(t, ipset.Adjacent(one_ten, one_one))
	assert.False(t, ipset.Adjacent(one_ten, two_hundred))
	assert.False(t, ipset.Adjacent(internet, two_hundred))
	assert.False(t, ipset.Adjacent(two_hundred, two_hundred))
	assert.False(t, ipset.Adjacent(internet, internet))
	assert.True(t, ipset.Adjacent(one_one, two_two))
	assert.True(t, ipset.Adjacent(one_one, two_hundred))
	assert.True(t, ipset.Adjacent(two_hundred, one_one))
	assert.True(t, ipset.Adjacent(hundredOne_hundredOne, two_hundred))
	assert.True(t, ipset.Adjacent(two_hundred, hundredOne_hundredOne))
}

func TestContiguous(t *testing.T) {
	f := ip.V6()
	min := ip.MinAddress(f)
	one := f.FromInt(1)
	two := f.FromInt(2)
	ten := f.FromInt(10)
	hundred := f.FromInt(100)
	hundredOne := f.FromInt(101)

	one_one := ipset.NewBlock(one, 128)
	two_two := ipset.NewBlock(two, 128)
	one_ten := ipset.NewInterval(one, ten)
	two_hundred := ipset.NewInterval(two, hundred)
	hundredOne_hundredOne := ipset.NewBlock(hundredOne, 128)
	internet := ipset.NewBlock(min, 0)

	assert.True(t, ipset.Contiguous(one_one, one_one))
	assert.True(t, ipset.Contiguous(one_one, one_ten))
	assert.True(t, ipset.Contiguous(one_ten, one_one))
	assert.True(t, ipset.Contiguous(one_ten, two_hundred))
	assert.True(t, ipset.Contiguous(internet, two_hundred))
	assert.True(t, ipset.Contiguous(two_hundred, two_hundred))
	assert.True(t, ipset.Contiguous(one_one, two_two))
	assert.True(t, ipset.Contiguous(one_one, two_hundred))
	assert.True(t, ipset.Contiguous(two_hundred, one_one))
	assert.True(t, ipset.Contiguous(hundredOne_hundredOne, two_hundred))
	assert.True(t, ipset.Contiguous(two_hundred, hundredOne_hundredOne))
	assert.False(t, ipset.Contiguous(two_two, hundredOne_hundredOne))
	assert.False(t, ipset.Contiguous(hundredOne_hundredOne, two_two))
}

func TestExtremes(t *testing.T) {
	f := ip.V6()
	min := ip.MinAddress(f)
	one := f.FromInt(1)
	hundredOne := f.FromInt(101)

	one_one := ipset.NewBlock(one, 128)
	hundredOne_hundredOne := ipset.NewBlock(hundredOne, 128)
	internet := ipset.NewBlock(min, 0)

	{
		actual := ipset.Extremes(one_one, hundredOne_hundredOne)
		assert.Equal(t, one, actual.First())
		assert.Equal(t, hundredOne, actual.Last())
	}
	{
		actual := ipset.Extremes(hundredOne_hundredOne, one_one)
		assert.Equal(t, one, actual.First())
		assert.Equal(t, hundredOne, actual.Last())
	}
	{
		actual := ipset.Extremes(internet, one_one)
		assert.Same(t, internet, actual)
	}
	{
		actual := ipset.Extremes(one_one, internet)
		assert.Same(t, internet, actual)
	}
}
