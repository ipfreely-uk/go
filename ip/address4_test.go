package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	one := ip.V4().FromInt(1)
	expected := ip.V4().FromInt(2)
	actual := one.Add(one)
	assert.Equal(t, expected, actual)
}

func TestSubtract(t *testing.T) {
	one := ip.V4().FromInt(1)
	expected := ip.V4().FromInt(0)
	actual := one.Subtract(one)
	assert.Equal(t, expected, actual)
}

func TestMultiply(t *testing.T) {
	two := ip.V4().FromInt(2)
	expected := ip.V4().FromInt(4)
	actual := two.Multiply(two)
	assert.Equal(t, expected, actual)
}

func TestDivide(t *testing.T) {
	two := ip.V4().FromInt(2)
	three := ip.V4().FromInt(3)
	expected := ip.V4().FromInt(1)
	actual := three.Divide(two)
	assert.Equal(t, expected, actual)
}

func TestMod(t *testing.T) {
	two := ip.V4().FromInt(2)
	three := ip.V4().FromInt(3)
	expected := ip.V4().FromInt(1)
	actual := three.Mod(two)
	assert.Equal(t, expected, actual)
}

func TestNot(t *testing.T) {
	zero := ip.V4().FromInt(0)
	expected := ip.V4().FromInt(0xFFFFFFFF)
	actual := zero.Not()
	assert.Equal(t, expected, actual)
}

func TestAnd(t *testing.T) {
	first := ip.V4().FromInt(0xAABBCC00)
	second := ip.V4().FromInt(0x00B000DD)
	expected := ip.V4().FromInt(0x00B00000)
	actual := first.And(second)
	assert.Equal(t, expected, actual)
}

func TestOr(t *testing.T) {
	first := ip.V4().FromInt(0xAABBCC00)
	second := ip.V4().FromInt(0x00B000DD)
	expected := ip.V4().FromInt(0xAABBCCDD)
	actual := first.Or(second)
	assert.Equal(t, expected, actual)
}

func TestXor(t *testing.T) {
	first := ip.V4().FromInt(0xAABBCC00)
	second := ip.V4().FromInt(0x00B000DD)
	expected := ip.V4().FromInt(0xAA0BCCDD)
	actual := first.Xor(second)
	assert.Equal(t, expected, actual)
}

func TestShiftRight(t *testing.T) {
	first := ip.V4().FromInt(0xAABBCC00)
	expected := ip.V4().FromInt(0x00AABBCC)
	actual := first.Shift(8)
	assert.Equal(t, expected, actual)
}

func TestShiftLeft(t *testing.T) {
	first := ip.V4().FromInt(0xAABBCC00)
	expected := ip.V4().FromInt(0xBBCC0000)
	actual := first.Shift(-8)
	assert.Equal(t, expected, actual)
}

func TestCompare(t *testing.T) {
	one := ip.V4().FromInt(1)
	hundred := ip.V4().FromInt(100)
	assert.Equal(t, -1, one.Compare(hundred))
	assert.Equal(t, 1, hundred.Compare(one))
	assert.Equal(t, 0, hundred.Compare(hundred))
}

func TestLeadingZeros(t *testing.T) {
	f := ip.V4()
	assert.Equal(t, 0, f.FromInt(0xFFFFFFFF).LeadingZeros())
	assert.Equal(t, 32, f.FromInt(0).LeadingZeros())
	assert.Equal(t, 8, f.FromInt(0xFFFFFF).LeadingZeros())
	assert.Equal(t, 8, f.FromInt(0xFFFF00).LeadingZeros())
	assert.Equal(t, 9, f.FromInt(0b01111111_11111111_11111111).LeadingZeros())
}

func TestTrailingZeros(t *testing.T) {
	f := ip.V4()
	assert.Equal(t, 0, f.FromInt(0xFFFFFFFF).TrailingZeros())
	assert.Equal(t, 32, f.FromInt(0).TrailingZeros())
	assert.Equal(t, 0, f.FromInt(0xFFFFFF).TrailingZeros())
	assert.Equal(t, 8, f.FromInt(0xFFFF00).TrailingZeros())
	assert.Equal(t, 1, f.FromInt(0b10).TrailingZeros())
}

func TestString(t *testing.T) {
	type test struct {
		input    []byte
		expected string
	}

	tests := []test{{
		[]byte{0, 0, 0, 0},
		"0.0.0.0",
	}, {
		[]byte{127, 0, 0, 1},
		"127.0.0.1",
	}, {
		[]byte{255, 255, 255, 255},
		"255.255.255.255",
	}}

	for _, candidate := range tests {
		a, err := ip.V4().FromBytes(candidate.input...)
		assert.Nil(t, err)
		actual := a.String()
		assert.Equal(t, candidate.expected, actual)
	}
}
