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
	v4 := ip.V4()
	r0 := exampleInterval(v4, "192.0.2.0", "192.0.2.100")
	r1 := exampleInterval(v4, "192.0.2.101", "192.0.2.111")
	r2 := exampleInterval(v4, "192.0.2.200", "192.0.2.200")

	union := ipset.NewDiscrete(r0, r1, r2)
	println(r0.String(), "\u222A", r1.String(), "\u222A", r2.String(), "=", union.String())
}

func exampleInterval[A ip.Number[A]](family ip.Family[A], first, last string) ipset.Interval[A] {
	a0 := ip.MustParse(family, first)
	a1 := ip.MustParse(family, last)
	return ipset.NewInterval(a0, a1)
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
