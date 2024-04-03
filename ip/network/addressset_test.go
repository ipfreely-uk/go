package network_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/stretchr/testify/assert"
)

func TestNewAddressSet(t *testing.T) {
	zero := ip.V6().FromInt(0)
	hundred := ip.V6().FromInt(100)
	thousand := ip.V6().FromInt(1000)
	r0 := network.NewRange(zero, zero)
	r1 := network.NewRange(hundred, thousand)
	r3 := network.NewRange(zero, thousand)
	set := network.NewSet(r0, r1, r3, r0)
	assert.NotNil(t, set)
}
