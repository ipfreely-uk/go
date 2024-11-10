package ip_test

import (
	"fmt"
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
	highest := lowest.Add(v4.FromInt(5))

	printOrder("Ascending", lowest, highest)
	for addr := range ip.Inclusive(lowest, highest) {
		println(addr.String())
	}

	printOrder("Descending", highest, lowest)
	for addr := range ip.Inclusive(highest, lowest) {
		println(addr.String())
	}
}

func printOrder(order string, start, end fmt.Stringer) {
	msg := fmt.Sprintf("%s: %s to %s", order, start.String(), end.String())
	println(msg)
}

func TestExampleExclusive(t *testing.T) {
	ExampleExclusive()
}

func ExampleExclusive() {
	v6 := ip.V6()
	docs := ip.MustParse(v6, "2001:DB8::")
	lowest := docs.Add(v6.FromInt(0xD))
	highest := lowest.Add(v6.FromInt(5))

	printOrder("Ascending", lowest, highest)
	for addr := range ip.Exclusive(lowest, highest) {
		println(addr.String())
	}

	printOrder("Descending", highest, lowest)
	for addr := range ip.Exclusive(highest, lowest) {
		println(addr.String())
	}
}
