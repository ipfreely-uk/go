// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset_test

import (
	"math/big"
	"math/rand"
	"slices"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipmask"
	. "github.com/ipfreely-uk/go/ipset"
	"github.com/stretchr/testify/assert"
)

func TestNewDiscrete(t *testing.T) {
	{
		zero := ip.V6().FromInt(0)
		hundred := ip.V6().FromInt(100)
		thousand := ip.V6().FromInt(1000)
		r0 := NewInterval(zero, zero)
		r1 := NewInterval(hundred, thousand)
		r3 := NewInterval(zero, thousand)
		set := NewDiscrete(r0, r1, r3, r0)
		assert.NotNil(t, set)
	}
	{
		empty := NewDiscrete[ip.Addr4]()
		empty = NewDiscrete(empty, empty, empty)
		assert.Equal(t, int64(0), empty.Size().Int64())
	}
	{
		zero := ip.V6().FromInt(0)
		two := ip.V6().FromInt(2)
		kay := ip.V6().FromInt(1024)
		expected := NewInterval(zero, kay)
		odds := []Discrete[ip.Addr6]{}
		for a := range expected.Addresses() {
			if !ip.Eq(a.Mod(two), zero) {
				odds = append(odds, NewSingle(a))
			}
		}
		evens := []Discrete[ip.Addr6]{}
		for a := range expected.Addresses() {
			if ip.Eq(a.Mod(two), zero) {
				evens = append(evens, NewSingle(a))
			}
		}
		// reverse evens
		slices.SortFunc(evens, func(a, b Discrete[ip.Addr6]) int {
			left := a.(Block[ip.Addr6])
			right := b.(Block[ip.Addr6])
			return right.First().Compare(left.First())
		})
		rogueIdx := 20
		rogue := evens[rogueIdx]
		evens = slices.Delete(evens, rogueIdx, rogueIdx)
		odd := NewDiscrete(odds...)
		even := NewDiscrete(evens...)
		actual := NewDiscrete(odd, even, rogue)
		assert.True(t, Eq(expected, actual))
	}
	{
		zero := ip.V6().FromInt(0)
		meg := ip.V6().FromInt(1024 * 1024)
		expected := NewInterval(zero, meg)
		contents := []Discrete[ip.Addr6]{}
		for a := range expected.Addresses() {
			contents = append(contents, NewSingle(a))
		}
		actual := NewDiscrete(contents...)
		assert.True(t, Eq(expected, actual))
	}
	{
		net, mask, err := ipmask.ParseCIDRNotation(ip.V4(), "10.0.0.0/24")
		assert.Nil(t, err)
		expected := NewBlock(net, mask)
		contents := []Discrete[ip.Addr4]{}
		for a := range expected.Addresses() {
			set := NewSingle(a)
			contents = append(contents, set)
		}
		src := rand.NewSource(0)
		ran := rand.New(src)
		ran.Shuffle(len(contents), func(i, j int) {
			left := contents[i]
			right := contents[j]
			contents[i] = right
			contents[j] = left
		})
		actual := NewDiscrete(contents...)
		assert.True(t, Eq(expected, actual))
	}
}

func TestDiscrete_Size(t *testing.T) {
	zero := ip.V6().FromInt(0)
	hundred := ip.V6().FromInt(100)
	thousand := ip.V6().FromInt(1000)
	r0 := NewInterval(zero, zero)
	r1 := NewInterval(hundred, thousand)
	set := NewDiscrete(r0, r1, r0)

	expected := big.NewInt(902)
	actual := set.Size()

	assert.Equal(t, expected, actual)
}

func TestDiscrete_Empty(t *testing.T) {
	{
		zero := ip.V6().FromInt(0)
		hundred := ip.V6().FromInt(100)
		thousand := ip.V6().FromInt(1000)
		r0 := NewInterval(zero, zero)
		r1 := NewInterval(hundred, thousand)
		set := NewDiscrete(r0, r1, r0)

		assert.False(t, set.Empty())
	}
	{
		empty := NewDiscrete[ip.Addr6]()
		d := NewDiscrete(empty, empty, empty)

		assert.True(t, d.Empty())
	}
}

func TestDiscrete_Contains(t *testing.T) {
	zero := ip.V6().FromInt(0)
	ninteynine := ip.V6().FromInt(99)
	hundred := ip.V6().FromInt(100)
	thousand := ip.V6().FromInt(1000)
	tenthousand := ip.V6().FromInt(10000)
	r0 := NewInterval(zero, zero)
	r1 := NewInterval(hundred, thousand)
	set := NewDiscrete(r0, r1, r0, r1)

	assert.True(t, set.Contains(zero))
	assert.False(t, set.Contains(ninteynine))
	assert.True(t, set.Contains(hundred))
	assert.True(t, set.Contains(thousand))
	assert.False(t, set.Contains(tenthousand))
}

func TestDiscrete_Addresses(t *testing.T) {
	zero := ip.V6().FromInt(0)
	hundred := ip.V6().FromInt(100)
	hundredAndOne := ip.V6().FromInt(101)
	r0 := NewInterval(zero, zero)
	r1 := NewInterval(hundred, hundredAndOne)
	set := NewDiscrete(r0, r1, r0, r1)

	{
		addresses := []ip.Addr6{}
		for a := range set.Addresses() {
			addresses = append(addresses, a)
		}

		assert.Equal(t, 3, len(addresses))
		assert.Equal(t, zero, addresses[0])
		assert.Equal(t, hundred, addresses[1])
		assert.Equal(t, hundredAndOne, addresses[2])
	}
	{
		addresses := []ip.Addr6{}
		for a := range set.Addresses() {
			addresses = append(addresses, a)
			break
		}

		assert.Equal(t, 1, len(addresses))
	}
}

func TestDiscrete_String(t *testing.T) {
	zero := ip.V6().FromInt(0)
	hundred := ip.V6().FromInt(100)
	thousand := ip.V6().FromInt(1000)
	r0 := NewInterval(zero, zero)
	r1 := NewInterval(hundred, thousand)
	set := NewDiscrete(r0, r1, r0, r1)

	actual := set.String()

	assert.Equal(t, "{::/128, ::64-::3e8}", actual)
}
