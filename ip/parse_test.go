// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"testing"

	. "github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	for _, c := range parsingTestSet() {
		if c.v == Version4 {
			actual, err := Parse(V4(), c.s)
			if err != nil {
				assert.False(t, c.ok)
				continue
			}
			expected := V4().MustFromBytes(c.b...)
			assert.Equal(t, expected, actual)
		} else {
			actual, err := Parse(V6(), c.s)
			if err != nil {
				assert.False(t, c.ok)
				continue
			}
			expected := V6().MustFromBytes(c.b...)
			assert.Equal(t, expected, actual)
		}
	}
}

func TestMustParse(t *testing.T) {
	for _, c := range parsingTestSet() {
		if c.v == Version4 {
			if c.ok {
				actual := MustParse(V4(), c.s)
				expected := V4().MustFromBytes(c.b...)
				assert.Equal(t, expected, actual)
			} else {
				assert.Panics(t, func() {
					MustParse(V4(), c.s)
				})
			}
		} else {
			if c.ok {
				actual := MustParse(V6(), c.s)
				expected := V6().MustFromBytes(c.b...)
				assert.Equal(t, expected, actual)
			} else {
				assert.Panics(t, func() {
					MustParse(V6(), c.s)
				})
			}
		}
	}
}

func TestParseUnknown(t *testing.T) {
	for _, c := range parsingTestSet() {
		actual, err := ParseUnknown(c.s)
		if err != nil {
			assert.False(t, c.ok, c.s)
			continue
		} else {
			assert.True(t, c.ok, c.s)
		}

		var expected any
		if c.v == Version4 {
			expected = V4().MustFromBytes(c.b...)
		} else {
			expected = V6().MustFromBytes(c.b...)
		}
		assert.Equal(t, expected, actual)
	}
}

func TestMustParseUnknown(t *testing.T) {
	for _, c := range parsingTestSet() {
		if c.ok {
			actual := MustParseUnknown(c.s)
			var expected any
			if c.v == Version4 {
				expected = V4().MustFromBytes(c.b...)
			} else {
				expected = V6().MustFromBytes(c.b...)
			}
			assert.Equal(t, expected, actual)
		} else {
			assert.Panics(t, func() { MustParseUnknown(c.s) })
		}
	}
}

func TestMustFromBytes(t *testing.T) {
	for _, c := range parsingTestSet() {
		if c.ok {
			actual := MustFromBytes(c.b...)
			var expected any
			if c.v == Version4 {
				expected = V4().MustFromBytes(c.b...)
			} else {
				expected = V6().MustFromBytes(c.b...)
			}
			assert.Equal(t, expected, actual)
		}
	}

	assert.Panics(t, func() { MustFromBytes() })
}

func TestFromBytes(t *testing.T) {
	for _, c := range parsingTestSet() {
		actual, err := FromBytes(c.b...)
		if err != nil {
			assert.False(t, c.ok)
			continue
		}
		if !c.ok {
			continue
		}
		var expected any
		if c.v == Version4 {
			expected = V4().MustFromBytes(c.b...)
		} else {
			expected = V6().MustFromBytes(c.b...)
		}
		assert.Equal(t, expected, actual)
	}
}

type parseTestCase struct {
	b  []byte
	s  string
	v  Version
	ok bool
}

func parsingTestSet() []parseTestCase {
	return []parseTestCase{
		{
			b:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			s:  "::",
			v:  Version6,
			ok: true,
		},
		{
			b:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			s:  "::1",
			v:  Version6,
			ok: true,
		},
		{
			b:  []byte{0xFE, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			s:  "fe80::1",
			v:  Version6,
			ok: true,
		}, {
			b:  []byte{0x20, 0x01, 0x0D, 0xB8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			s:  "2001:db8::",
			v:  Version6,
			ok: true,
		},
		{
			b:  []byte{0xFF, 0xFF, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xCA, 0xFE, 0xBA, 0xBE},
			s:  "FFFF::Cafe:Babe",
			v:  Version6,
			ok: true,
		},
		{
			b:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			s:  "1:2:3:4:5:6:7:8:9",
			v:  Version6,
			ok: false,
		},
		{
			b:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			s:  "1:2:3:4:5::6:7:8",
			v:  Version6,
			ok: false,
		},
		{
			b:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			s:  "::1:2:3:4:5:6:7:8",
			v:  Version6,
			ok: false,
		},
		{
			b:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			s:  "1:2:3:4:5:6:7:8::",
			v:  Version6,
			ok: false,
		},
		{
			b:  []byte{0, 0, 0, 0},
			s:  "0:0:0:0",
			v:  Version6,
			ok: false,
		},
		{
			b:  []byte{},
			s:  ":",
			v:  Version6,
			ok: false,
		},
		{
			b:  []byte{},
			s:  ":::",
			v:  Version6,
			ok: false,
		},
		{
			b:  []byte{},
			s:  "10000::",
			v:  Version6,
			ok: false,
		},
		{
			b:  []byte{0, 0, 0},
			s:  ":::",
			v:  Version6,
			ok: false,
		},
		{
			b:  []byte{0, 0, 0, 0},
			s:  "0.0.0.0",
			v:  Version4,
			ok: true,
		},
		{
			b:  []byte{0, 0, 0, 1},
			s:  "0.0.0.1",
			v:  Version4,
			ok: true,
		},
		{
			b:  []byte{255, 255, 255, 255},
			s:  "255.255.255.255",
			v:  Version4,
			ok: true,
		},
		{
			b:  []byte{255, 255, 255},
			s:  "255.255.255.",
			v:  Version4,
			ok: false,
		},
		{
			b:  []byte{255, 255, 255},
			s:  "0255.255.255.255",
			v:  Version4,
			ok: false,
		},
		{
			b:  []byte{255, 255, 255},
			s:  "256.255.255.255",
			v:  Version4,
			ok: false,
		},
		{
			b:  []byte{255, 255, 255},
			s:  "255..255.255",
			v:  Version4,
			ok: false,
		},
		{
			b:  []byte{255, 255, 255},
			s:  ".255.255.255",
			v:  Version4,
			ok: false,
		},
		{
			b:  []byte{1, 0, 0, 0},
			s:  "01.0.0.0", // decimal or octal ambiguity
			v:  Version4,
			ok: false,
		},
		{
			b:  []byte{1, 0, 0, 0},
			s:  "0.0.0.008", // decimal or octal ambiguity
			v:  Version4,
			ok: false,
		},
		{
			b:  []byte{1, 0, 0, 0},
			s:  "0.0.0.0255", // decimal or octal ambiguity
			v:  Version4,
			ok: false,
		},
		{
			b:  []byte{0, 0, 0, 0},
			s:  "000.0.0.0",
			v:  Version4,
			ok: true,
		},
	}
}
