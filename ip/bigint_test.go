// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestBigInt(t *testing.T) {
	two := big.NewInt(2)
	expected := ip.V4().FromInt(2)
	actual, err := ip.FromBigInt(ip.V4(), two)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	actual2 := ip.ToBigInt(expected)
	assert.Equal(t, two, actual2)

	max := ip.ToBigInt(ip.MaxAddress(ip.V4()))
	oversize := max.Mul(max, two)
	_, err = ip.FromBigInt(ip.V4(), oversize)
	assert.NotNil(t, err)

	max = ip.ToBigInt(ip.MaxAddress(ip.V4()))
	back, err := ip.FromBigInt(ip.V4(), max)
	assert.Nil(t, err)
	assert.Equal(t, ip.MaxAddress(ip.V4()), back)
}
