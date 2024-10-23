package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/ipset"
	"github.com/stretchr/testify/assert"
)

func TestSeq(t *testing.T) {
	netaddr, mask, _ := ipset.ParseCIDRNotation(ip.V4(), "192.0.2.128/28")
	subnet := ipset.NewBlock(netaddr, mask)
	{
		count := 0
		for _ = range ipset.Seq(subnet.Addresses()) {
			count++
		}
		assert.Equal(t, 16, count)
	}
	{
		s := ipset.Seq(subnet.Addresses())
		s(func(a ip.Addr4) bool {
			return false
		})
	}
}
