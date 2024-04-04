package cidr_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/ipfreely-uk/go/ip/network/cidr"
	"github.com/stretchr/testify/assert"
)

func TestNotation(t *testing.T) {
	a := ip.V6().FromInt(0)
	mask := ip.V6().Width()
	block := network.NewBlock(a, mask)
	actual := cidr.Notation(block)
	assert.Equal(t, "::/128", actual)
}

func TestParseUnknown(t *testing.T) {
	legal := []string{
		"192.168.0.0/24",
		"192.168.0.0/32",
		"::/0",
		"::/128",
	}
	for _, c := range legal {
		b, err := cidr.ParseUnknown(c)
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
		_, err := cidr.ParseUnknown(c)
		assert.NotNil(t, err)
	}
}
