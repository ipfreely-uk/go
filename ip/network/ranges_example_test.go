package network_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
)

func TestExampleContiguous(t *testing.T) {
	ExampleContiguous()
}

func ExampleContiguous() {
	r0 := makeRange(ip.V6(), "2001:db8::", "2001:db8::100")
	r1 := makeRange(ip.V6(), "2001:db8::10", "2001:db8::ffff:ffff:ffff")

	if network.Contiguous(r0, r1) {
		r2 := network.Join(r0, r1)
		println(r2.String())
	}
}

func makeRange[A ip.Number[A]](family ip.Family[A], first, last string) network.AddressRange[A] {
	f := ip.MustParse(family, first)
	l := ip.MustParse(family, last)
	return network.NewRange(f, l)
}
