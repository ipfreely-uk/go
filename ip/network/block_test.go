package network_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/ipfreely-uk/go/ip/subnet"
	"github.com/stretchr/testify/assert"
)

func TestNewBlock(t *testing.T) {
	address, _ := ip.V4().FromBytes(192, 168, 0, 0)
	subnet := network.NewBlock(address, 24)
	assert.NotNil(t, subnet)

	assert.Panics(t, func() {
		network.NewBlock(address, 0)
	})
}

func TestBlock_MaskSize(t *testing.T) {
	address, _ := ip.V4().FromBytes(192, 168, 0, 0)
	mask := network.NewBlock(address, 24).MaskSize()
	assert.Equal(t, 24, mask)
}

func TestBlock_Size(t *testing.T) {
	address, _ := ip.V4().FromBytes(192, 168, 0, 0)
	actual := network.NewBlock(address, 24).Size()
	expected := subnet.AddressCount(ip.V4(), 24)
	assert.Equal(t, expected, actual)
}

func TestBlock_Contains(t *testing.T) {
	address, _ := ip.V4().FromBytes(192, 168, 0, 0)
	actual := network.NewBlock(address, 24)

	assert.True(t, actual.Contains(address))
	assert.False(t, actual.Contains(ip.MaxAddress(ip.V4())))
}

func TestBlock_Addresses(t *testing.T) {
	address, _ := ip.V4().FromBytes(192, 168, 0, 0)
	actual := network.NewBlock(address, 24)

	count := big.NewInt(0)
	one := big.NewInt(1)
	next := actual.Addresses()
	for _, exists := next(); exists; _, exists = next() {
		count = count.Add(count, one)
	}
	assert.Equal(t, actual.Size(), count)
}

func TestBlock_Ranges(t *testing.T) {
	address, _ := ip.V4().FromBytes(192, 168, 0, 0)
	actual := network.NewBlock(address, 24)

	count := big.NewInt(0)
	one := big.NewInt(1)
	next := actual.Ranges()
	for _, exists := next(); exists; _, exists = next() {
		count = count.Add(count, one)
	}
	assert.Equal(t, one, count)
}

func TestBlock_Mask(t *testing.T) {
	address, _ := ip.V4().FromBytes(192, 168, 0, 0)
	actual := network.NewBlock(address, 24).Mask()
	expected := subnet.Mask(ip.V4(), 24)
	assert.Equal(t, expected, actual)
}
