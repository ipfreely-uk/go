package ipset_test

import (
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	empty := ipset.NewDiscrete[ip.Addr4]()

	assert.Equal(t, big.NewInt(0), empty.Size())
	assert.False(t, empty.Contains(ip.MaxAddress(ip.V4())))
	assert.Equal(t, "{}", empty.String())
	{
		count := 0
		for _ = range empty.Addresses() {
			count++
		}
		assert.Equal(t, 0, count)
	}
	{
		count := 0
		for _ = range empty.Intervals() {
			count++
		}
		assert.Equal(t, 0, count)
	}
}
