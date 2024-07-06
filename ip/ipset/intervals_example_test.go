package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/ipset"
)

func TestExampleJoin(t *testing.T) {
	ExampleJoin()
}

func ExampleJoin() {
	set := func(a0, a1 string) ipset.Interval[ip.Addr6] {
		v6 := ip.V6()
		p := ip.MustParse[ip.Addr6]
		return ipset.NewInterval(p(v6, a0), p(v6, a1))
	}
	r0 := set("2001:db8::", "2001:db8::100")
	r1 := set("2001:db8::10", "2001:db8::ffff:ffff:ffff")

	if ipset.Contiguous(r0, r1) {
		r2 := ipset.Join(r0, r1)
		println(r2.String())
	}
}
