package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"

	"github.com/stretchr/testify/assert"
)

func TestAdd6(t *testing.T) {
	one := ip.V6().FromInt(1)
	expected := ip.V6().FromInt(2)
	actual := one.Add(one)
	assert.Equal(t, expected, actual)
}

func TestSubtract6(t *testing.T) {
	one := ip.V6().FromInt(1)
	expected := ip.V6().FromInt(0)
	actual := one.Subtract(one)
	assert.Equal(t, expected, actual)
}

func TestMultiply6(t *testing.T) {
	two := ip.V6().FromInt(2)
	expected := ip.V6().FromInt(4)
	actual := two.Multiply(two)
	assert.Equal(t, expected, actual)
}

func TestDivide6(t *testing.T) {
	two := ip.V6().FromInt(2)
	three := ip.V6().FromInt(3)
	expected := ip.V6().FromInt(1)
	actual := three.Divide(two)
	assert.Equal(t, expected, actual)
}

func TestMod6(t *testing.T) {
	two := ip.V6().FromInt(2)
	three := ip.V6().FromInt(3)
	expected := ip.V6().FromInt(1)
	actual := three.Mod(two)
	assert.Equal(t, expected, actual)
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
	assert.Equal(t, -1, one.Compare(hundred))
	assert.Equal(t, 1, hundred.Compare(one))
	assert.Equal(t, 0, hundred.Compare(hundred))
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
