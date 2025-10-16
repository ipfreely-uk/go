// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"testing"

	. "github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestWidth6(t *testing.T) {
	assert.Equal(t, 128, V6().Width())
}

func TestVersion6(t *testing.T) {
	assert.Equal(t, Version6, V6().Version())
}

func TestFromInt6(t *testing.T) {
	address := V6().FromInt(0xFFFFFFFF)
	bytes := address.Bytes()
	for i := range bytes {
		var expected byte = 0xFF
		if i < 12 {
			expected = 0
		}
		assert.Equal(t, expected, bytes[i])
	}
}

func TestFromBytes6(t *testing.T) {
	expected := V6().FromInt(0xCCAAFFEE)
	actual, err := V6().FromBytes(0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xCC, 0xAA, 0xFF, 0xEE)
	bs := actual.Bytes()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
	assert.Equal(t, byte(0), bs[0])
	assert.Equal(t, byte(0), bs[11])
	assert.Equal(t, byte(0xCC), bs[12])
	assert.Equal(t, byte(0xAA), bs[13])
	assert.Equal(t, byte(0xFF), bs[14])
	assert.Equal(t, byte(0xEE), bs[15])

	assert.Equal(t, actual, V6().MustFromBytes(actual.Bytes()...))

	_, err = V6().FromBytes()
	assert.NotNil(t, err)

	assert.Panics(t, func() { V6().MustFromBytes() })
}

func TestFromInvalidBytes6(t *testing.T) {
	_, err := V4().FromBytes(0xFF, 0xFF, 0xFF)
	assert.NotNil(t, err)
}

func TestFamilyString6(t *testing.T) {
	s := V6().String()
	assert.Equal(t, "IPv6", s)
}
