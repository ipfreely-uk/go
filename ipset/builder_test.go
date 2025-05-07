package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	zero := ip.V4().FromInt(0)
	ffff := ip.V4().FromInt(0xffff)
	expected := ipset.NewInterval(zero, ffff)

	bld := ipset.NewBuilder[ip.Addr4]().Threshold(64)
	for a := range expected.Addresses() {
		single := ipset.NewSingle(a)
		bld.Or(single)
	}
	actual := bld.Union()

	assert.True(t, ipset.Eq(expected, actual))
}
