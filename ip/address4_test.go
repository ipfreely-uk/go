// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ipfreely-uk/go/ip"

	"github.com/stretchr/testify/assert"
)

func TestMaths4(t *testing.T) {
	v4 := ip.V4()

	values := []ip.Addr4{
		ip.MinAddress(v4),
		ip.MaxAddress(v4),
		ip.MustParse(v4, "255.0.0.0"),
		ip.MustParse(v4, "0.0.0.255"),
		ip.MustParse(v4, "0.0.0.1"),
		ip.MustParse(v4, "0.0.0.2"),
	}

	bytes := make([]byte, v4.Width()/8)
	ran := rand.New(rand.NewSource(0))
	for i := 0; i < 200; i++ {
		_, err := ran.Read(bytes)
		if err != nil {
			t.Fail()
			return
		}

		a := v4.MustFromBytes(bytes...)
		values = append(values, a)
	}

	testMaths(t, values)
}

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

func TestA4_Float64(t *testing.T) {
	min := ip.MinAddress(ip.V4())
	max := min.Not()
	two := ip.V4().FromInt(2)
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

func TestA4_Format(t *testing.T) {
	{
		a := ip.MaxAddress(ip.V4())
		expected := a.String()
		actual := fmt.Sprintf("%v", a)
		assert.Equal(t, expected, actual)
	}
	{
		a := ip.MaxAddress(ip.V4())
		expected := a.String()
		actual := fmt.Sprintf("%s", a)
		assert.Equal(t, expected, actual)
	}
	{
		a := ip.MaxAddress(ip.V4())
		expected := ip.ToBigInt(a).String()
		actual := fmt.Sprintf("%d", a)
		assert.Equal(t, expected, actual)
	}
	{
		a := ip.MaxAddress(ip.V4())
		expected := fmt.Sprintf("%x", ip.ToBigInt(a))
		actual := fmt.Sprintf("%x", a)
		assert.Equal(t, expected, actual)
	}
	{
		a := ip.MaxAddress(ip.V4())
		expected := fmt.Sprintf("%X", ip.ToBigInt(a))
		actual := fmt.Sprintf("%X", a)
		assert.Equal(t, expected, actual)
	}
	{
		a := ip.MaxAddress(ip.V4())
		expected := fmt.Sprintf("%b", ip.ToBigInt(a))
		actual := fmt.Sprintf("%b", a)
		assert.Equal(t, expected, actual)
	}
}
