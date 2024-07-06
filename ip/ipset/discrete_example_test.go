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
	set := func(first, last string) ipset.Interval[ip.Addr4] {
		v4 := ip.V4()
		p := ip.MustParse[ip.Addr4]
		return ipset.NewInterval(p(v4, first), p(v4, last))
	}
	r0 := set("192.0.2.0", "192.0.2.100")
	r1 := set("192.0.2.101", "192.0.2.111")
	r2 := set("192.0.2.200", "192.0.2.200")

	union := ipset.NewDiscrete(r0, r1, r2)
	println(r0.String(), "\u222A", r1.String(), "\u222A", r2.String(), "=", union.String())
}

func TestExampleNewDiscrete_second(t *testing.T) {
	ExampleNewDiscrete_second()
}

func ExampleNewDiscrete_second() {
	printEmptySetFor(ip.V4())
	printEmptySetFor(ip.V6())
}

func printEmptySetFor[A ip.Number[A]](f ip.Family[A]) {
	empty := ipset.NewDiscrete[A]()
	println(f.String(), empty.String())
}
