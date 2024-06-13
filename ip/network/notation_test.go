package network_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip/network"
	"github.com/stretchr/testify/assert"
)

func TestParseUnknown(t *testing.T) {
	legal := []string{
		"192.168.0.0/24",
		"192.168.0.0/32",
		"::/0",
		"::/128",
	}
	for _, c := range legal {
		b, err := network.ParseUnknownCIDRNotation(c)
		assert.Nil(t, err)
		assert.NotNil(t, b)
	}
	illegal := []string{
		"foo/24",
		"192.168.0.0/128",
		"192.168.0.0/0",
		"192.168.0.0/",
		"192.168.0.0",
		"::/-1",
		"::/129",
	}
	for _, c := range illegal {
		_, err := network.ParseUnknownCIDRNotation(c)
		assert.NotNil(t, err)
	}
}
