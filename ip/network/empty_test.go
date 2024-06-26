package network_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	empty := network.NewSet[ip.Addr4]()

	assert.Equal(t, big.NewInt(0), empty.Size())
	assert.False(t, empty.Contains(ip.MaxAddress(ip.V4())))
	_, exists := empty.Addresses()()
	assert.False(t, exists)
	_, exists = empty.Ranges()()
	assert.False(t, exists)
	assert.Equal(t, "{}", empty.String())
}
