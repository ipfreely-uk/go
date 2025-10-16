// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipmask_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	. "github.com/ipfreely-uk/go/ipmask"
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

	assert.Panics(t, func() { For(ip.V4(), 33) })
	assert.Panics(t, func() { For(ip.V6(), 129) })
	assert.Panics(t, func() { For(ip.V4(), -1) })
}

func verifyMask[A ip.Int[A]](t *testing.T, expected []byte, family ip.Family[A], mask int) {
	e, x := family.FromBytes(expected...)
	assert.Nil(t, x)
	actual := For(family, mask)
	assert.Equal(t, e, actual)
}

func TestAddressCount(t *testing.T) {
	one := big.NewInt(1)
	v6, _ := big.NewInt(0).SetString("340282366920938463463374607431768211456", 10)

	assert.Equal(t, one, Size(ip.V4(), 32))
	assert.Equal(t, v6, Size(ip.V6(), 0))

	assert.Panics(t, func() { Size(ip.V4(), 33) })
	assert.Panics(t, func() { Size(ip.V6(), 129) })
	assert.Panics(t, func() { Size(ip.V6(), -1) })
}

func TestMaskSize(t *testing.T) {
	{
		one := ip.V4().FromInt(1)
		two := ip.V4().FromInt(2)
		tenStart, _ := ip.Parse(ip.V4(), "10.0.0.0")
		tenEnd, _ := ip.Parse(ip.V4(), "10.255.255.255")

		assert.Equal(t, 32, Test(one, one))
		assert.Equal(t, 8, Test(tenStart, tenEnd))
		assert.Equal(t, 0, Test(ip.MinAddress(ip.V4()), ip.MaxAddress(ip.V4())))
		assert.Equal(t, -1, Test(one, tenEnd))
		assert.Equal(t, -1, Test(one, two))
	}
	{
		one := ip.V6().FromInt(1)
		two := ip.V6().FromInt(2)
		fe80Start, _ := ip.Parse(ip.V6(), "fe80::")
		fe80End, _ := ip.Parse(ip.V6(), "fe80:ffff:ffff:ffff:ffff:ffff:ffff:ffff")

		assert.Equal(t, 128, Test(one, one))
		assert.Equal(t, 16, Test(fe80Start, fe80End))
		assert.Equal(t, 0, Test(ip.MinAddress(ip.V6()), ip.MaxAddress(ip.V6())))
		assert.Equal(t, -1, Test(one, fe80End))
		assert.Equal(t, -1, Test(one, two))
	}
}

func TestSubnetCovers(t *testing.T) {
	fe80 := ip.MustParse(ip.V6(), "fe80::")
	zero := ip.V6().FromInt(0)

	assert.True(t, Covers(128, fe80))
	assert.True(t, Covers(16, fe80))
	assert.True(t, Covers(0, zero))

	assert.False(t, Covers(0, fe80))
	assert.False(t, Covers(1, fe80))
	assert.False(t, Covers(8, fe80))

	assert.Panics(t, func() {
		Covers(-1, fe80)
	})
	assert.Panics(t, func() {
		Covers(129, fe80)
	})
}

func TestSubnetMaskBits(t *testing.T) {
	{
		invalid := ip.MustParse(ip.V6(), "f0ff::")
		actual := Bits(invalid)
		assert.Equal(t, -1, actual, invalid.String())
	}
	{
		v4 := ip.V4()
		for i := 0; i < ip.Width4; i++ {
			m := For(v4, i)
			actual := Bits(m)
			assert.Equal(t, i, actual, m.String())
		}
	}
	{
		v6 := ip.V6()
		for i := 0; i < ip.Width6; i++ {
			m := For(v6, i)
			actual := Bits(m)
			assert.Equal(t, i, actual, m.String())
		}
	}
}
