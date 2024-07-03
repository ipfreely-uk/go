package ip_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestMask(t *testing.T) {
	verifyMask(t, []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, ip.V6(), 0)
	verifyMask(t, []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, ip.V6(), 128)
	verifyMask(t, []byte{0xFF, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, ip.V6(), 8)
	verifyMask(t, []byte{0xFF, 0, 0, 0}, ip.V4(), 8)
	verifyMask(t, []byte{0, 0, 0, 0}, ip.V4(), 0)
	verifyMask(t, []byte{0xFF, 0xFF, 0xFF, 0xFF}, ip.V4(), 32)
	verifyMask(t, []byte{0b10000000, 0, 0, 0}, ip.V4(), 1)
	verifyMask(t, []byte{0b11000000, 0, 0, 0}, ip.V4(), 2)
	verifyMask(t, []byte{0b11100000, 0, 0, 0}, ip.V4(), 3)
	verifyMask(t, []byte{0b11110000, 0, 0, 0}, ip.V4(), 4)
	verifyMask(t, []byte{0b11111000, 0, 0, 0}, ip.V4(), 5)
	verifyMask(t, []byte{0b11111100, 0, 0, 0}, ip.V4(), 6)
	verifyMask(t, []byte{0b11111110, 0, 0, 0}, ip.V4(), 7)

	assert.Panics(t, func() { ip.SubnetMask(ip.V4(), 33) })
	assert.Panics(t, func() { ip.SubnetMask(ip.V6(), 129) })
	assert.Panics(t, func() { ip.SubnetMask(ip.V4(), -1) })
}

func verifyMask[A ip.Number[A]](t *testing.T, expected []byte, family ip.Family[A], mask int) {
	e, x := family.FromBytes(expected...)
	assert.Nil(t, x)
	actual := ip.SubnetMask(family, mask)
	assert.Equal(t, e, actual)
}

func TestAddressCount(t *testing.T) {
	one := big.NewInt(1)
	v6, _ := big.NewInt(0).SetString("340282366920938463463374607431768211456", 10)

	assert.Equal(t, one, ip.SubnetAddressCount(ip.V4(), 32))
	assert.Equal(t, v6, ip.SubnetAddressCount(ip.V6(), 0))

	assert.Panics(t, func() { ip.SubnetAddressCount(ip.V4(), 33) })
	assert.Panics(t, func() { ip.SubnetAddressCount(ip.V6(), 129) })
	assert.Panics(t, func() { ip.SubnetAddressCount(ip.V6(), -1) })
}

func TestMaskSize(t *testing.T) {
	{
		one := ip.V4().FromInt(1)
		two := ip.V4().FromInt(2)
		tenStart, _ := ip.Parse(ip.V4(), "10.0.0.0")
		tenEnd, _ := ip.Parse(ip.V4(), "10.255.255.255")

		assert.Equal(t, 32, ip.SubnetMaskSize(one, one))
		assert.Equal(t, 8, ip.SubnetMaskSize(tenStart, tenEnd))
		assert.Equal(t, 0, ip.SubnetMaskSize(ip.MinAddress(ip.V4()), ip.MaxAddress(ip.V4())))
		assert.Equal(t, -1, ip.SubnetMaskSize(one, tenEnd))
		assert.Equal(t, -1, ip.SubnetMaskSize(one, two))
	}
	{
		one := ip.V6().FromInt(1)
		two := ip.V6().FromInt(2)
		fe80Start, _ := ip.Parse(ip.V6(), "fe80::")
		fe80End, _ := ip.Parse(ip.V6(), "fe80:ffff:ffff:ffff:ffff:ffff:ffff:ffff")

		assert.Equal(t, 128, ip.SubnetMaskSize(one, one))
		assert.Equal(t, 16, ip.SubnetMaskSize(fe80Start, fe80End))
		assert.Equal(t, 0, ip.SubnetMaskSize(ip.MinAddress(ip.V6()), ip.MaxAddress(ip.V6())))
		assert.Equal(t, -1, ip.SubnetMaskSize(one, fe80End))
		assert.Equal(t, -1, ip.SubnetMaskSize(one, two))
	}
}

func TestSubnetCovers(t *testing.T) {
	fe80 := ip.MustParse(ip.V6(), "fe80::")
	zero := ip.V6().FromInt(0)

	assert.True(t, ip.SubnetMaskCovers(128, fe80))
	assert.True(t, ip.SubnetMaskCovers(16, fe80))
	assert.True(t, ip.SubnetMaskCovers(0, zero))

	assert.False(t, ip.SubnetMaskCovers(-1, fe80))
	assert.False(t, ip.SubnetMaskCovers(0, fe80))
	assert.False(t, ip.SubnetMaskCovers(129, fe80))
}
