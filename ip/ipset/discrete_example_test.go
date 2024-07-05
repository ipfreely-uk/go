package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/ipset"
)

func TestExampleNewDiscrete(t *testing.T) {
	ExampleNewDiscrete()
}

func ExampleNewDiscrete() {
	family := ip.V4()
	r0 := exampleInterval(family, "192.0.2.0", "192.0.2.100")
	r1 := exampleInterval(family, "192.0.2.101", "192.0.2.111")
	r2 := exampleInterval(family, "192.0.2.200", "192.0.2.255")
	r3 := exampleInterval(family, "203.0.113.0", "203.0.113.255")
	r4 := exampleInterval(family, "192.0.2.0", "192.0.2.100")

	addresses := ipset.NewDiscrete(r0, r1, r2, r3, r4)

	println("Rationalized ranges:", addresses.String())
}

func exampleInterval[A ip.Number[A]](family ip.Family[A], first, last string) ipset.Interval[A] {
	a0 := ip.MustParse(family, first)
	a1 := ip.MustParse(family, last)
	return ipset.NewInterval(a0, a1)
}
