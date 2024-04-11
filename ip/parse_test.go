package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	for _, c := range parsingTestSet() {
		if c.v == ip.Version4 {
			actual, err := ip.Parse(ip.V4(), c.s)
			if err != nil {
				assert.False(t, c.ok)
				continue
			}
			expected := ip.V4().MustFromBytes(c.b...)
			assert.Equal(t, expected, actual)
		} else {
			actual, err := ip.Parse(ip.V6(), c.s)
			if err != nil {
				assert.False(t, c.ok)
				continue
			}
			expected := ip.V6().MustFromBytes(c.b...)
			assert.Equal(t, expected, actual)
		}
	}
}

func TestMustParse(t *testing.T) {
	for _, c := range parsingTestSet() {
		if c.v == ip.Version4 {
			if c.ok {
				actual := ip.MustParse(ip.V4(), c.s)
				expected := ip.V4().MustFromBytes(c.b...)
				assert.Equal(t, expected, actual)
			} else {
				assert.Panics(t, func() {
					ip.MustParse(ip.V4(), c.s)
				})
			}
		} else {
			if c.ok {
				actual := ip.MustParse(ip.V6(), c.s)
				expected := ip.V6().MustFromBytes(c.b...)
				assert.Equal(t, expected, actual)
			} else {
				assert.Panics(t, func() {
					ip.MustParse(ip.V6(), c.s)
				})
			}
		}
	}
}

func TestParseUnknown(t *testing.T) {
	for _, c := range parsingTestSet() {
		actual, err := ip.ParseUnknown(c.s)
		if err != nil {
			assert.False(t, c.ok)
			continue
		}

		var expected any
		if c.v == ip.Version4 {
			expected = ip.V4().MustFromBytes(c.b...)
		} else {
			expected = ip.V6().MustFromBytes(c.b...)
		}
		assert.Equal(t, expected, actual)
	}
}

func TestMustParseUnknown(t *testing.T) {
	for _, c := range parsingTestSet() {
		if c.ok {
			actual := ip.MustParseUnknown(c.s)
			var expected any
			if c.v == ip.Version4 {
				expected = ip.V4().MustFromBytes(c.b...)
			} else {
				expected = ip.V6().MustFromBytes(c.b...)
			}
			assert.Equal(t, expected, actual)
		} else {
			assert.Panics(t, func() { ip.MustParseUnknown(c.s) })
		}
	}
}

func TestMustFromBytes(t *testing.T) {
	for _, c := range parsingTestSet() {
		if c.ok {
			actual := ip.MustFromBytes(c.b...)
			var expected any
			if c.v == ip.Version4 {
				expected = ip.V4().MustFromBytes(c.b...)
			} else {
				expected = ip.V6().MustFromBytes(c.b...)
			}
			assert.Equal(t, expected, actual)
		}
	}

	assert.Panics(t, func() { ip.MustFromBytes() })
}

func TestFromBytes(t *testing.T) {
	for _, c := range parsingTestSet() {
		actual, err := ip.FromBytes(c.b...)
		if err != nil {
			assert.False(t, c.ok)
			continue
		}
		if !c.ok {
			continue
		}
		var expected any
		if c.v == ip.Version4 {
			expected = ip.V4().MustFromBytes(c.b...)
		} else {
			expected = ip.V6().MustFromBytes(c.b...)
		}
		assert.Equal(t, expected, actual)
	}
}

type parseTestCase struct {
	b  []byte
	s  string
	v  ip.Version
	ok bool
}

func parsingTestSet() []parseTestCase {
	return []parseTestCase{
		{
			b:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			s:  "::",
			v:  ip.Version6,
			ok: true,
		},
		{
			b:  []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			s:  "::1",
			v:  ip.Version6,
			ok: true,
		},
		{
			b:  []byte{0xFE, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			s:  "fe80::1",
			v:  ip.Version6,
			ok: true,
		},
		{
			b:  []byte{0xFF, 0xFF, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0xCA, 0xFE, 0xBA, 0xBE},
			s:  "FFFF::Cafe:Babe",
			v:  ip.Version6,
			ok: true,
		},
		{
			b:  []byte{0, 0, 0, 0},
			s:  "0:0:0:0",
			v:  ip.Version6,
			ok: false,
		},
		{
			b:  []byte{},
			s:  ":",
			v:  ip.Version6,
			ok: false,
		},
		{
			b:  []byte{},
			s:  "10000::",
			v:  ip.Version6,
			ok: false,
		},
		{
			b:  []byte{0, 0, 0},
			s:  ":::",
			v:  ip.Version6,
			ok: false,
		},
		{
			b:  []byte{0, 0, 0, 0},
			s:  "0.0.0.0",
			v:  ip.Version4,
			ok: true,
		},
		{
			b:  []byte{0, 0, 0, 1},
			s:  "0.0.0.1",
			v:  ip.Version4,
			ok: true,
		},
		{
			b:  []byte{255, 255, 255, 255},
			s:  "255.255.255.255",
			v:  ip.Version4,
			ok: true,
		},
		{
			b:  []byte{255, 255, 255},
			s:  "255.255.255.",
			v:  ip.Version4,
			ok: false,
		},
		{
			b:  []byte{255, 255, 255},
			s:  "0255.255.255.255",
			v:  ip.Version4,
			ok: false,
		},
		{
			b:  []byte{255, 255, 255},
			s:  "256.255.255.255",
			v:  ip.Version4,
			ok: false,
		},
	}
}
