package network_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/stretchr/testify/assert"
)

func TestNewRange(t *testing.T) {
	first := ip.V6().FromInt(1)
	last := ip.V6().FromInt(10)

	actual := network.NewRange(first, last)

	assert.Equal(t, first, actual.First())
	assert.Equal(t, last, actual.Last())
}

func TestRange_Contains(t *testing.T) {
	zero := ip.V6().FromInt(0)
	one := ip.V6().FromInt(1)
	two := ip.V6().FromInt(2)
	three := ip.V6().FromInt(3)
	four := ip.V6().FromInt(4)

	actual := network.NewRange(one, three)

	assert.True(t, actual.Contains(one))
	assert.True(t, actual.Contains(two))
	assert.True(t, actual.Contains(three))
	assert.False(t, actual.Contains(zero))
	assert.False(t, actual.Contains(four))
}

func TestRange_Size(t *testing.T) {
	one := ip.V6().FromInt(1)
	three := ip.V6().FromInt(3)

	expected := big.NewInt(3)
	actual := network.NewRange(one, three).Size()

	assert.Equal(t, expected, actual)
}

func TestRange_Addresses(t *testing.T) {
	one := ip.V6().FromInt(1)
	three := ip.V6().FromInt(3)

	actual := network.NewRange(one, three).Addresses()

	var count = 0
	var last ip.Address6
	for ok, addr := actual(); ok; ok, addr = actual() {
		last = addr
		count++
	}
	assert.Equal(t, three, last)
	assert.Equal(t, 3, count)
}

func TestRange_Ranges(t *testing.T) {
	one := ip.V6().FromInt(1)
	three := ip.V6().FromInt(3)
	net := network.NewRange(one, three)

	actual := net.Ranges()

	var count = 0
	var last network.AddressRange[ip.Address6]
	for ok, addr := actual(); ok; ok, addr = actual() {
		last = addr
		count++
	}
	assert.Same(t, net, last)
	assert.Equal(t, 1, count)
}
