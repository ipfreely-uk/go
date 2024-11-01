package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
)

func TestExampleNext(t *testing.T) {
	ExampleNext()
}

func ExampleNext() {
	v4 := ip.V4()
	lowest := v4.MustFromBytes(192, 0, 2, 0)
	highest := lowest.Add(v4.FromInt(10))

	Ascend(lowest, highest)
}

func Ascend[A ip.Int[A]](lowest, highest A) {
	current := lowest
	for {
		println(current.String())
		if ip.Eq(current, highest) {
			break
		}
		current = ip.Next(current)
	}
}

func TestExamplePrev(t *testing.T) {
	ExamplePrev()
}

func ExamplePrev() {
	v6 := ip.V6()
	highest := ip.MustParse(v6, "2001:db8::fffe")
	lowest := highest.Subtract(v6.FromInt(10))

	Descend(highest, lowest)
}

func Descend[A ip.Int[A]](highest, lowest A) {
	current := highest
	for {
		println(current.String())
		if ip.Eq(current, lowest) {
			break
		}
		current = ip.Prev(current)
	}
}

func TestExampleInclusive(t *testing.T) {
	ExampleInclusive()
}

func ExampleInclusive() {
	v4 := ip.V4()
	lowest := v4.MustFromBytes(192, 0, 2, 0)
	highest := lowest.Add(v4.FromInt(10))

	for addr := range ip.Inclusive(highest, lowest) {
		println(addr.String())
	}
}

func ExampleInclusive_second() {
	v4 := ip.V4()
	lowest := v4.MustFromBytes(192, 0, 2, 0)
	highest := lowest.Add(v4.FromInt(10))

	for addr := range ip.Inclusive(lowest, highest) {
		println(addr.String())
	}
}

func TestExampleExclusive(t *testing.T) {
	ExampleExclusive()
}

func ExampleExclusive() {
	v4 := ip.V4()
	lowest := v4.MustFromBytes(192, 0, 2, 0)
	highest := lowest.Add(v4.FromInt(10))

	for addr := range ip.Exclusive(lowest, highest) {
		println(addr.String())
	}
}
