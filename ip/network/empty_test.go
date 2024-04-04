package network_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	empty := network.NewSet[ip.Address4]()

	assert.Equal(t, big.NewInt(0), empty.Size())
	assert.False(t, empty.Contains(ip.MaxAddress(ip.V4())))
	ok, _ := empty.Addresses()()
	assert.False(t, ok)
	ok, _ = empty.Ranges()()
	assert.False(t, ok)
}
