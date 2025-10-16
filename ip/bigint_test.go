// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"math/big"
	"testing"

	. "github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestBigInt(t *testing.T) {
	two := big.NewInt(2)
	expected := V4().FromInt(2)
	actual, err := FromBigInt(V4(), two)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)

	actual2 := ToBigInt(expected)
	assert.Equal(t, two, actual2)

	max := ToBigInt(MaxAddress(V4()))
	oversize := max.Mul(max, two)
	_, err = FromBigInt(V4(), oversize)
	assert.NotNil(t, err)

	max = ToBigInt(MaxAddress(V4()))
	back, err := FromBigInt(V4(), max)
	assert.Nil(t, err)
	assert.Equal(t, MaxAddress(V4()), back)
}
