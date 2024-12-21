package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestMinAddress(t *testing.T) {
	expected := ip.V6().FromInt(0)
	actual := ip.MinAddress(ip.V6())
	assert.Equal(t, expected, actual)
}

func TestMaxAddress(t *testing.T) {
	expected := ip.V4().FromInt(0xFFFFFFFF)
	actual := ip.MaxAddress(ip.V4())
	assert.Equal(t, expected, actual)
}

func TestNext(t *testing.T) {
	expected := ip.V4().FromInt(0)
	max := ip.V4().FromInt(0xFFFFFFFF)
	actual := ip.Next(max)
	assert.Equal(t, expected, actual)
}

func TestPrev(t *testing.T) {
	expected := ip.V6().FromInt(0)
	one := ip.V6().FromInt(1)
	actual := ip.Prev(one)
	assert.Equal(t, expected, actual)
}

func TestInclusive(t *testing.T) {
	smallest := ip.V6().FromInt(0)
	middle := ip.Next(smallest)
	largest := ip.Next(middle)
	{
		count := 0
		for range ip.Inclusive(smallest, largest) {
			count++
		}
		assert.Equal(t, 3, count)
	}
	{
		count := 0
		for range ip.Inclusive(largest, smallest) {
			count++
		}
		assert.Equal(t, 3, count)
	}
	{
		count := 0
		for range ip.Inclusive(smallest, smallest) {
			count++
		}
		assert.Equal(t, 1, count)
	}
	{
		yield := func(a ip.Addr6) bool {
			return false
		}
		ip.Inclusive(largest, smallest)(yield)
	}
}

func TestExclusive(t *testing.T) {
	smallest := ip.V6().FromInt(0)
	middle := ip.Next(smallest)
	largest := ip.Next(middle)
	{
		count := 0
		for range ip.Exclusive(smallest, largest) {
			count++
		}
		assert.Equal(t, 2, count)
	}
	{
		count := 0
		for range ip.Exclusive(largest, smallest) {
			count++
		}
		assert.Equal(t, 2, count)
	}
	{
		count := 0
		for range ip.Exclusive(smallest, smallest) {
			count++
		}
		assert.Equal(t, 0, count)
	}
	{
		yield := func(a ip.Addr6) bool {
			return false
		}
		ip.Exclusive(largest, smallest)(yield)
	}
}
