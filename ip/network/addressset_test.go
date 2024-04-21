package network_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/stretchr/testify/assert"
)

func TestNewAddressSet(t *testing.T) {
	zero := ip.V6().FromInt(0)
	hundred := ip.V6().FromInt(100)
	thousand := ip.V6().FromInt(1000)
	r0 := network.NewRange(zero, zero)
	r1 := network.NewRange(hundred, thousand)
	r3 := network.NewRange(zero, thousand)
	set := network.NewSet(r0, r1, r3, r0)
	assert.NotNil(t, set)
}

func TestAddressSet_Size(t *testing.T) {
	zero := ip.V6().FromInt(0)
	hundred := ip.V6().FromInt(100)
	thousand := ip.V6().FromInt(1000)
	r0 := network.NewRange(zero, zero)
	r1 := network.NewRange(hundred, thousand)
	set := network.NewSet(r0, r1, r0)

	expected := big.NewInt(902)
	actual := set.Size()

	assert.Equal(t, expected, actual)
}

func TestAddressSet_Contains(t *testing.T) {
	zero := ip.V6().FromInt(0)
	hundred := ip.V6().FromInt(100)
	thousand := ip.V6().FromInt(1000)
	tenthousand := ip.V6().FromInt(10000)
	r0 := network.NewRange(zero, zero)
	r1 := network.NewRange(hundred, thousand)
	set := network.NewSet(r0, r1, r0, r1)

	assert.True(t, set.Contains(zero))
	assert.True(t, set.Contains(hundred))
	assert.True(t, set.Contains(thousand))
	assert.False(t, set.Contains(tenthousand))
}

func TestAddressSet_Addresses(t *testing.T) {
	zero := ip.V6().FromInt(0)
	hundred := ip.V6().FromInt(100)
	hundredAndOne := ip.V6().FromInt(101)
	r0 := network.NewRange(zero, zero)
	r1 := network.NewRange(hundred, hundredAndOne)
	set := network.NewSet(r0, r1, r0, r1)

	addresses := []ip.Addr6{}
	iter := set.Addresses()
	for a, exists := iter(); exists; a, exists = iter() {
		addresses = append(addresses, a)
	}

	assert.Equal(t, 3, len(addresses))
	assert.Equal(t, zero, addresses[0])
	assert.Equal(t, hundred, addresses[1])
	assert.Equal(t, hundredAndOne, addresses[2])
}

func TestAddressSet_String(t *testing.T) {
	zero := ip.V6().FromInt(0)
	hundred := ip.V6().FromInt(100)
	thousand := ip.V6().FromInt(1000)
	r0 := network.NewRange(zero, zero)
	r1 := network.NewRange(hundred, thousand)
	set := network.NewSet(r0, r1, r0, r1)

	actual := set.String()

	assert.Equal(t, "{::/128, ::64-::3e8}", actual)
}
