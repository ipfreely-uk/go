package subnet_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/subnet"
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
}

func verifyMask[A ip.Address[A]](t *testing.T, expected []byte, family ip.Family[A], mask int) {
	e, x := family.FromBytes(expected...)
	assert.Nil(t, x)
	actual := subnet.Mask(family, mask)
	assert.Equal(t, e, actual)
}

func TestAddressCount(t *testing.T) {
	one := big.NewInt(1)
	v6, _ := big.NewInt(0).SetString("340282366920938463463374607431768211456", 10)

	assert.Equal(t, one, subnet.AddressCount(ip.V4(), 32))
	assert.Equal(t, v6, subnet.AddressCount(ip.V6(), 0))
}

func TestMaskSize(t *testing.T) {
	one := ip.V4().FromInt(1)
	tenStart, _ := ip.Parse(ip.V4(), "10.0.0.0")
	tenEnd, _ := ip.Parse(ip.V4(), "10.255.255.255")

	assert.Equal(t, 32, subnet.MaskSize(one, one))
	assert.Equal(t, 8, subnet.MaskSize(tenStart, tenEnd))
	assert.Equal(t, 0, subnet.MaskSize(ip.MinAddress(ip.V4()), ip.MaxAddress(ip.V4())))
	assert.Equal(t, 0, subnet.MaskSize(ip.MinAddress(ip.V6()), ip.MaxAddress(ip.V6())))
	assert.Equal(t, -1, subnet.MaskSize(one, tenEnd))
}
