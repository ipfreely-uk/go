package ip_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"

	"github.com/stretchr/testify/assert"
)

func TestAdd6(t *testing.T) {
	one := ip.V6().FromInt(1)
	expected := ip.V6().FromInt(2)
	actual := one.Add(one)
	assert.Equal(t, expected, actual)

	big, _ := ip.Parse(ip.V6(), "::ffff:ffff:ffff:ffff")
	expected, _ = ip.Parse(ip.V6(), "0:0:0:1::")
	actual = big.Add(one)
	assert.Equal(t, expected, actual)
}

func TestSubtract6(t *testing.T) {
	one := ip.V6().FromInt(1)
	expected := ip.V6().FromInt(0)
	actual := one.Subtract(one)
	assert.Equal(t, expected, actual)

	big, _ := ip.Parse(ip.V6(), "fe80:0:0:1::")
	expected, _ = ip.Parse(ip.V6(), "fe80:0:0:0:ffff:ffff:ffff:ffff")
	actual = big.Subtract(one)
	assert.Equal(t, expected, actual)
}

func TestMultiply6(t *testing.T) {
	zero := ip.V6().FromInt(0)
	one := ip.V6().FromInt(1)
	two := ip.V6().FromInt(2)
	four := ip.V6().FromInt(4)

	actual := two.Multiply(two)
	assert.Equal(t, four, actual)

	actual = four.Multiply(one)
	assert.Equal(t, four, actual)

	actual = one.Multiply(four)
	assert.Equal(t, four, actual)

	actual = four.Multiply(zero)
	assert.Equal(t, zero, actual)

	actual = zero.Multiply(four)
	assert.Equal(t, zero, actual)
}

func TestDivide6(t *testing.T) {
	zero := ip.V6().FromInt(0)
	one := ip.V6().FromInt(1)
	two := ip.V6().FromInt(2)
	three := ip.V6().FromInt(3)

	actual := three.Divide(two)
	assert.Equal(t, one, actual)

	actual = three.Divide(one)
	assert.Equal(t, three, actual)

	actual = one.Divide(three)
	assert.Equal(t, zero, actual)

	actual = three.Divide(three)
	assert.Equal(t, one, actual)

	assert.Panics(t, func() { zero.Divide(zero) })
}

func TestMod6(t *testing.T) {
	zero := ip.V6().FromInt(0)
	one := ip.V6().FromInt(1)
	two := ip.V6().FromInt(2)
	three := ip.V6().FromInt(3)

	actual := three.Mod(two)
	assert.Equal(t, one, actual)

	actual = three.Mod(three)
	assert.Equal(t, zero, actual)

	assert.Panics(t, func() { zero.Mod(zero) })
}

func TestNot6(t *testing.T) {
	zero := ip.V6().FromInt(0)
	expected, _ := ip.V6().FromBytes(0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF)
	actual := zero.Not()
	assert.Equal(t, expected, actual)
}

func TestAnd6(t *testing.T) {
	first := ip.V6().FromInt(0xAABBCC00)
	second := ip.V6().FromInt(0x00B000DD)
	expected := ip.V6().FromInt(0x00B00000)
	actual := first.And(second)
	assert.Equal(t, expected, actual)
}

func TestOr6(t *testing.T) {
	first := ip.V6().FromInt(0xAABBCC00)
	second := ip.V6().FromInt(0x00B000DD)
	expected := ip.V6().FromInt(0xAABBCCDD)
	actual := first.Or(second)
	assert.Equal(t, expected, actual)
}

func TestXor6(t *testing.T) {
	first := ip.V6().FromInt(0xAABBCC00)
	second := ip.V6().FromInt(0x00B000DD)
	expected := ip.V6().FromInt(0xAA0BCCDD)
	actual := first.Xor(second)
	assert.Equal(t, expected, actual)
}

func TestShiftRight6(t *testing.T) {
	first, _ := ip.V6().FromBytes(0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x11)
	expected, _ := ip.V6().FromBytes(0x00, 0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF)
	actual := first.Shift(8)
	assert.Equal(t, expected, actual)
}

func TestShiftLeft6(t *testing.T) {
	first, _ := ip.V6().FromBytes(0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x11)
	expected, _ := ip.V6().FromBytes(0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, 0x99, 0xAA, 0xBB, 0xCC, 0xDD, 0xEE, 0xFF, 0x11, 0x00)
	actual := first.Shift(-8)
	assert.Equal(t, expected, actual)
}

func TestCompare6(t *testing.T) {
	one := ip.V6().FromInt(1)
	hundred := ip.V6().FromInt(100)
	max := ip.V6().FromInt(0).Not()
	assert.Equal(t, -1, one.Compare(hundred))
	assert.Equal(t, 1, hundred.Compare(one))
	assert.Equal(t, 0, hundred.Compare(hundred))
	assert.Equal(t, 1, max.Compare(one))
	assert.Equal(t, -1, one.Compare(max))
}

func TestLeadingZeros6(t *testing.T) {
	f := ip.V6()
	two := f.FromInt(2)
	assert.Equal(t, 0, f.FromInt(0).Not().LeadingZeros())
	assert.Equal(t, 128, f.FromInt(0).LeadingZeros())
	assert.Equal(t, 104, f.FromInt(0xFF_FF_FF).LeadingZeros())
	assert.Equal(t, 104, f.FromInt(0xFF_FF_00).LeadingZeros())
	assert.Equal(t, 105, f.FromInt(0b01111111_11111111_11111111).LeadingZeros())
	assert.Equal(t, 0, f.FromInt(1).Not().LeadingZeros())
	assert.Equal(t, 1, ip.MaxAddress(f).Divide(two).LeadingZeros())
	bottom, _ := ip.Parse(ip.V6(), "::FFFF:FFFF:FFFF:FFFF")
	assert.Equal(t, 64, bottom.LeadingZeros())
}

func TestTrailingZeros6(t *testing.T) {
	f := ip.V6()
	assert.Equal(t, 0, f.FromInt(0xFFFFFFFF).TrailingZeros())
	assert.Equal(t, 128, f.FromInt(0).TrailingZeros())
	assert.Equal(t, 0, f.FromInt(0xFFFFFF).TrailingZeros())
	assert.Equal(t, 8, f.FromInt(0xFFFF00).TrailingZeros())
	assert.Equal(t, 1, f.FromInt(0b10).TrailingZeros())
	top, _ := ip.Parse(ip.V6(), "FFFF:FFFF:FFFF:FFFF::")
	assert.Equal(t, 64, top.TrailingZeros())
}

func TestString6(t *testing.T) {
	type test struct {
		input    []byte
		expected string
	}

	tests := []test{{
		[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		"::",
	}, {
		[]byte{0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xca, 0xfe},
		"ff00::cafe",
	}, {
		[]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xf},
		"1:203:405:607:809:a0b:c0d:e0f",
	}}

	for _, candidate := range tests {
		a, err := ip.V6().FromBytes(candidate.input...)
		assert.Nil(t, err)
		actual := a.String()
		assert.Equal(t, candidate.expected, actual)
	}
}

func TestA6_Float64(t *testing.T) {
	min := ip.MinAddress(ip.V6())
	max := min.Not()
	two := ip.V6().FromInt(2)
	half := max.Divide(two)

	expected, _ := big.NewInt(0).Float64()
	actual := min.Float64()
	assert.Equal(t, expected, actual)

	expected, _ = ip.ToBigInt(max).Float64()
	actual = max.Float64()
	assert.Equal(t, expected, actual)

	expected, _ = ip.ToBigInt(two).Float64()
	actual = two.Float64()
	assert.Equal(t, expected, actual)

	expected, _ = ip.ToBigInt(half).Float64()
	actual = half.Float64()
	assert.Equal(t, expected, actual)
}
