package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	zero := ip.V6().FromInt(0)
	megs := ip.V6().FromInt(10 * 1024 * 1024)
	expected := ipset.NewInterval(zero, megs)

	b := ipset.Builder[ip.Addr6]{}
	for addr := range expected.Addresses() {
		single := ipset.NewSingle(addr)
		b.Union(single)
	}
	actual := b.Build()

	assert.True(t, ipset.Eq(expected, actual))
}
