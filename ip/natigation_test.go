// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"testing"

	. "github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestMinAddress(t *testing.T) {
	expected := V6().FromInt(0)
	actual := MinAddress(V6())
	assert.Equal(t, expected, actual)
}

func TestMaxAddress(t *testing.T) {
	expected := V4().FromInt(0xFFFFFFFF)
	actual := MaxAddress(V4())
	assert.Equal(t, expected, actual)
}

func TestNext(t *testing.T) {
	expected := V4().FromInt(0)
	max := V4().FromInt(0xFFFFFFFF)
	actual := Next(max)
	assert.Equal(t, expected, actual)
}

func TestPrev(t *testing.T) {
	expected := V6().FromInt(0)
	one := V6().FromInt(1)
	actual := Prev(one)
	assert.Equal(t, expected, actual)
}

func TestInclusive(t *testing.T) {
	smallest := V6().FromInt(0)
	middle := Next(smallest)
	largest := Next(middle)
	{
		count := 0
		for range Inclusive(smallest, largest) {
			count++
		}
		assert.Equal(t, 3, count)
	}
	{
		count := 0
		for range Inclusive(largest, smallest) {
			count++
		}
		assert.Equal(t, 3, count)
	}
	{
		count := 0
		for range Inclusive(smallest, smallest) {
			count++
		}
		assert.Equal(t, 1, count)
	}
	{
		yield := func(a Addr6) bool {
			return false
		}
		Inclusive(largest, smallest)(yield)
	}
}

func TestExclusive(t *testing.T) {
	smallest := V6().FromInt(0)
	middle := Next(smallest)
	largest := Next(middle)
	{
		count := 0
		for range Exclusive(smallest, largest) {
			count++
		}
		assert.Equal(t, 2, count)
	}
	{
		count := 0
		for range Exclusive(largest, smallest) {
			count++
		}
		assert.Equal(t, 2, count)
	}
	{
		count := 0
		for range Exclusive(smallest, smallest) {
			count++
		}
		assert.Equal(t, 0, count)
	}
	{
		yield := func(a Addr6) bool {
			return false
		}
		Exclusive(largest, smallest)(yield)
	}
}
