package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
)

func TestExampleExtremes(t *testing.T) {
	ExampleExtremes()
}

func ExampleExtremes() {
	r0 := parseV6Interval("2001:db8::", "2001:db8::100")
	r1 := parseV6Interval("2001:db8::10", "2001:db8::ffff:ffff:ffff")

	if ipset.Contiguous(r0, r1) {
		r2 := ipset.Extremes(r0, r1)
		println(r2.String())
	}
}

func parseV6Interval(first, last string) ipset.Interval[ip.Addr6] {
	v6 := ip.V6()
	p := ip.MustParse[ip.Addr6]
	return ipset.NewInterval(p(v6, first), p(v6, last))
}
