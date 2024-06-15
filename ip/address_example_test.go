package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
)

func TestExampleAddress(t *testing.T) {
	ExampleAddress()
}

func ExampleAddress() {
	printN(ip.V4().MustFromBytes(192, 0, 2, 0), 3)
	printN(ip.MustParse(ip.V6(), "2001:db8::"), 3)
}

func printN[A ip.Address[A]](address A, n int) {
	a := address
	for i := 0; i < n; i++ {
		println(a.String())
		a = ip.Next(a)
	}
}

func TestExampleUntyped(t *testing.T) {
	ExampleUntyped()
}

func ExampleUntyped() {
	examples := []string{
		"2001:db8::",
		"192.0.2.0",
	}

	for _, e := range examples {
		address := ip.MustParseUnknown(e)
		switch a := address.(type) {
		case ip.Addr4:
			printNthAfter(a, 255)
		case ip.Addr6:
			printNthAfter(a, 0xFFFFFFFF)
		}
	}
}

func printNthAfter[A ip.Address[A]](address A, n uint32) {
	operand := address.Family().FromInt(n)
	result := address.Add(operand)
	println(result.String())
}
