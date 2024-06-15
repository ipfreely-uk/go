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
	first := v4.MustFromBytes(192, 0, 2, 0)
	last := first.Add(v4.FromInt(255))

	Ascend(first, last)
}

func Ascend[A ip.Address[A]](lowest, highest A) {
	current := lowest
	for {
		println(current.String())
		if current.Compare(highest) == 0 {
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
	first := ip.MustParse(v6, "2001:db8::")
	last := first.Add(v6.FromInt(255))

	Descend(last, first)
}

func Descend[A ip.Address[A]](highest, lowest A) {
	current := highest
	for {
		println(current.String())
		if current.Compare(lowest) == 0 {
			break
		}
		current = ip.Prev(current)
	}
}
