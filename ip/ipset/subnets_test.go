package ipset_test

import (
	"iter"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/ipset"
	"github.com/stretchr/testify/assert"
)

func TestSubnets(t *testing.T) {
	netAddr := ip.MustParse(ip.V4(), "192.168.0.0")
	{
		b := ipset.NewBlock(netAddr, 31)
		next, stop := iter.Pull(ipset.Subnets(b, 32))
		defer stop()

		b, exists := next()
		assert.True(t, exists)
		assert.Equal(t, "192.168.0.0/32", b.String())

		b, exists = next()
		assert.True(t, exists)
		assert.Equal(t, "192.168.0.1/32", b.String())

		_, exists = next()
		assert.False(t, exists)
	}
	{
		b := ipset.NewBlock(netAddr, 31)
		assert.Panics(t, func() {
			ipset.Subnets(b, 12)
		})
		assert.Panics(t, func() {
			ipset.Subnets(b, 33)
		})
	}
}
