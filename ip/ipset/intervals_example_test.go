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
	r0 := makeInterval(ip.V6(), "2001:db8::", "2001:db8::100")
	r1 := makeInterval(ip.V6(), "2001:db8::10", "2001:db8::ffff:ffff:ffff")

	if ipset.Contiguous(r0, r1) {
		r2 := ipset.Join(r0, r1)
		println(r2.String())
	}
}

func makeInterval[A ip.Number[A]](family ip.Family[A], first, last string) ipset.Interval[A] {
	f := ip.MustParse(family, first)
	l := ip.MustParse(family, last)
	return ipset.NewInterval(f, l)
}
