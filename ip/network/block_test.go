package network_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/stretchr/testify/assert"
)

func TestNewBlock(t *testing.T) {
	{
		address, _ := ip.V4().FromBytes(192, 168, 0, 0)
		subnet := network.NewBlock(address, 24)
		assert.NotNil(t, subnet)

		assert.Panics(t, func() {
			network.NewBlock(address, 0)
		})
	}
	{
		address := ip.MustParse(ip.V6(), "fe80::")
		block := network.NewBlock(address, 128)
		assert.NotNil(t, block)
	}
}

func TestBlock_MaskSize(t *testing.T) {
	{
		address, _ := ip.V4().FromBytes(192, 168, 0, 0)
		mask := network.NewBlock(address, 24).MaskSize()
		assert.Equal(t, 24, mask)
	}
	{
		address := ip.MustParse(ip.V6(), "fe80::")
		block := network.NewBlock(address, 128)
		assert.Equal(t, 128, block.MaskSize())
	}
}

func TestBlock_Size(t *testing.T) {
	{
		address, _ := ip.V4().FromBytes(192, 168, 0, 0)
		actual := network.NewBlock(address, 24).Size()
		expected := ip.SubnetAddressCount(ip.V4(), 24)
		assert.Equal(t, expected, actual)
	}
	{
		address := ip.MustParse(ip.V6(), "fe80::")
		actual := network.NewBlock(address, 128).Size()
		expected := ip.SubnetAddressCount(ip.V6(), 128)
		assert.Equal(t, expected, actual)
	}
}

func TestBlock_Contains(t *testing.T) {
	{
		address, _ := ip.V4().FromBytes(192, 168, 0, 0)
		actual := network.NewBlock(address, 24)
		assert.True(t, actual.Contains(address))
		assert.False(t, actual.Contains(ip.MaxAddress(ip.V4())))
	}
	{
		address := ip.MustParse(ip.V6(), "fe80::")
		actual := network.NewBlock(address, 128)
		assert.True(t, actual.Contains(address))
		assert.False(t, actual.Contains(ip.MinAddress(ip.V6())))
	}
}

func TestBlock_Addresses(t *testing.T) {
	{
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
	{
		address := ip.MustParse(ip.V6(), "fe80::")
		actual := network.NewBlock(address, 128)

		count := big.NewInt(0)
		one := big.NewInt(1)
		next := actual.Addresses()
		for _, exists := next(); exists; _, exists = next() {
			count = count.Add(count, one)
		}
		assert.Equal(t, actual.Size(), count)
	}
}

func TestBlock_Ranges(t *testing.T) {
	{
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
	{
		address := ip.MustParse(ip.V6(), "fe80::")
		actual := network.NewBlock(address, 128)

		count := big.NewInt(0)
		one := big.NewInt(1)
		next := actual.Ranges()
		for _, exists := next(); exists; _, exists = next() {
			count = count.Add(count, one)
		}
		assert.Equal(t, one, count)
	}
}

func TestBlock_Mask(t *testing.T) {
	{
		address, _ := ip.V4().FromBytes(192, 168, 0, 0)
		actual := network.NewBlock(address, 24).Mask()
		expected := ip.SubnetMask(ip.V4(), 24)
		assert.Equal(t, expected, actual)
	}
	{
		address := ip.MustParse(ip.V6(), "fe80::")
		actual := network.NewBlock(address, 128).Mask()
		expected := ip.SubnetMask(ip.V6(), 128)
		assert.Equal(t, expected, actual)
	}
}

func TestBlock_CidrNotation(t *testing.T) {
	{
		address, _ := ip.V4().FromBytes(192, 168, 0, 0)
		actual := network.NewBlock(address, 24).CidrNotation()
		expected := "192.168.0.0/24"
		assert.Equal(t, expected, actual)
	}
	{
		address := ip.MustParse(ip.V6(), "fe80::")
		actual := network.NewBlock(address, 128).CidrNotation()
		expected := "fe80::/128"
		assert.Equal(t, expected, actual)
	}
}
