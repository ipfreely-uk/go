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

	combined := union(r0, r1)
	println(r0.String(), "\u222A", r1.String(), "=", combined.String())
}

// Example union function
func union[A ip.Number[A]](sets ...ipset.Discrete[A]) ipset.Discrete[A] {
	slice := []ipset.Interval[A]{}
	for _, set := range sets {
		slice = appendToSlice(set.Intervals(), slice)
	}
	return ipset.NewDiscrete(slice...)
}

// Example iterator to slice function
func appendToSlice[E any](i ipset.Iterator[E], slice []E) []E {
	result := slice
	for e, ok := i(); ok; e, ok = i() {
		result = append(result, e)
	}
	return result
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
	empty := emptySet[A]()
	println(f.Version(), empty.String())
}

// Example empty set function
func emptySet[A ip.Number[A]]() ipset.Discrete[A] {
	return ipset.NewDiscrete[A]()
}
