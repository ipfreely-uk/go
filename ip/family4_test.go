// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestWidth4(t *testing.T) {
	assert.Equal(t, 32, ip.V4().Width())
}

func TestVersion4(t *testing.T) {
	assert.Equal(t, ip.Version4, ip.V4().Version())
}

func TestFromInt4(t *testing.T) {
	address := ip.V4().FromInt(0xFFFFFFFF)
	bytes := address.Bytes()
	for i := range bytes {
		var expected byte = 0xFF
		assert.Equal(t, expected, bytes[i])
	}
}

func TestFromBytes4(t *testing.T) {
	expected := ip.V4().FromInt(0xCCAAFFEE)
	actual, err := ip.V4().FromBytes(0xCC, 0xAA, 0xFF, 0xEE)
	bs := actual.Bytes()
	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
	assert.Equal(t, byte(0xCC), bs[0])
	assert.Equal(t, byte(0xAA), bs[1])
	assert.Equal(t, byte(0xFF), bs[2])
	assert.Equal(t, byte(0xEE), bs[3])

	assert.Equal(t, actual, ip.V4().MustFromBytes(actual.Bytes()...))

	_, err = ip.V4().FromBytes()
	assert.NotNil(t, err)

	assert.Panics(t, func() { ip.V4().MustFromBytes() })
}

func TestFromInvalidBytes4(t *testing.T) {
	_, err := ip.V4().FromBytes(0xFF, 0xFF, 0xFF)
	assert.NotNil(t, err)
}

func TestFamilyString4(t *testing.T) {
	s := ip.V4().String()
	assert.Equal(t, "IPv4", s)
}
