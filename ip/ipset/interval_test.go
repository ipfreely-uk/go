package ipset_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/ipset"
	"github.com/stretchr/testify/assert"
)

func TestNewInterval(t *testing.T) {
	first := ip.V6().FromInt(1)
	last := ip.V6().FromInt(10)

	actual := ipset.NewInterval(first, last)

	assert.Equal(t, first, actual.First())
	assert.Equal(t, last, actual.Last())
}

func TestInterval_Contains(t *testing.T) {
	zero := ip.V6().FromInt(0)
	one := ip.V6().FromInt(1)
	two := ip.V6().FromInt(2)
	three := ip.V6().FromInt(3)
	four := ip.V6().FromInt(4)

	actual := ipset.NewInterval(one, three)

	assert.True(t, actual.Contains(one))
	assert.True(t, actual.Contains(two))
	assert.True(t, actual.Contains(three))
	assert.False(t, actual.Contains(zero))
	assert.False(t, actual.Contains(four))
}

func TestInterval_Size(t *testing.T) {
	one := ip.V6().FromInt(1)
	three := ip.V6().FromInt(3)

	expected := big.NewInt(3)
	actual := ipset.NewInterval(one, three).Size()

	assert.Equal(t, expected, actual)
}

func TestInterval_Addresses(t *testing.T) {
	one := ip.V6().FromInt(1)
	three := ip.V6().FromInt(3)

	actual := ipset.NewInterval(one, three)

	var count = 0
	var last ip.Addr6
	for addr := range actual.Addresses() {
		last = addr
		count++
	}
	assert.Equal(t, three, last)
	assert.Equal(t, 3, count)
}

func TestInterval_Intervals(t *testing.T) {
	one := ip.V6().FromInt(1)
	three := ip.V6().FromInt(3)
	net := ipset.NewInterval(one, three)

	var count = 0
	var last ipset.Interval[ip.Addr6]
	for addr := range net.Intervals() {
		last = addr
		count++
	}
	assert.Same(t, net, last)
	assert.Equal(t, 1, count)
}
