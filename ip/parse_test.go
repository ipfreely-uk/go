package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestParse4(t *testing.T) {
	type test struct {
		expected []byte
		input    string
	}

	tests := []test{
		{
			[]byte{127, 0, 0, 1},
			"127.0.0.1",
		},
		{
			[]byte{255, 255, 255, 255},
			"255.255.255.255",
		},
		{
			[]byte{0, 0, 0, 0},
			"0.0.0.0",
		},
	}

	for _, candidate := range tests {
		expected, x := ip.V4().FromBytes(candidate.expected...)
		assert.Nil(t, x)
		actual, err := ip.Parse(ip.V4(), candidate.input)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}
}

func TestParseBad4(t *testing.T) {
	tests := []string{
		"127.0.0.1a",
		"127.0.0",
		"127.0..0.1",
		"::",
	}

	for _, candidate := range tests {
		_, err := ip.Parse(ip.V4(), candidate)
		assert.NotNil(t, err)
	}
}

func TestParse6(t *testing.T) {
	type test struct {
		expected []byte
		input    string
	}

	tests := []test{
		{
			[]byte{0xFE, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			"fe80::1",
		},
		{
			[]byte{0xFE, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			"FE80::1",
		},
		{
			[]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			"::",
		},
	}

	for _, c := range tests {
		expected, x := ip.V6().FromBytes(c.expected...)
		assert.Nil(t, x)
		actual, err := ip.Parse(ip.V6(), c.input)
		assert.Nil(t, err)
		assert.Equal(t, expected, actual)
	}
}

func TestParseBad6(t *testing.T) {
	tests := []string{
		"127.0.0.1a",
		":",
		":::",
		"fe800::",
		"foobar",
	}

	for _, candidate := range tests {
		_, err := ip.Parse(ip.V4(), candidate)
		assert.NotNil(t, err)
	}
}

func TestParseUnknown(t *testing.T) {
	type test struct {
		expected []byte
		input    string
		err      bool
	}

	tests := []test{
		{
			expected: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			input:    "::",
			err:      false,
		},
		{
			expected: []byte{0, 0, 0, 0},
			input:    "foo",
			err:      true,
		},
		{
			expected: []byte{0, 0, 0, 0},
			input:    "0.0.0.0",
			err:      false,
		},
	}

	for _, c := range tests {
		expected, _ := ip.FromBytes(c.expected...)
		actual, err := ip.ParseUnknown(c.input)
		if c.err {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, expected, actual)
		}
	}
}
