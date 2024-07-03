package ip_test

import (
	"net/netip"
	"testing"

	"github.com/ipfreely-uk/go/ip"
)

func TestExampleFromBytes(t *testing.T) {
	ExampleFromBytes()
}

func ExampleFromBytes() {
	address, err := ip.FromBytes(192, 0, 2, 1)
	if err != nil {
		println("Not address:", err)
	}
	switch a := address.(type) {
	case ip.Addr4:
		println("IPv4 address:", a.String())
	case ip.Addr6:
		println("IPv6 address", a.String())
	}
}

func TestExampleMustFromBytes(t *testing.T) {
	ExampleMustFromBytes()
}

func ExampleMustFromBytes() {
	// Convert to/from netip.Addr

	examples := []string{
		"2001:db8::",
		"192.0.2.1",
	}

	for _, e := range examples {
		original := netip.MustParseAddr(e)

		addr := fromNetip(original)

		var result netip.Addr
		switch a := addr.(type) {
		case ip.Addr4:
			result = toNetip(a)
		case ip.Addr6:
			result = toNetip(a)
		}

		println(original.String(), "->", result.String())
	}
}

func toNetip[A ip.Number[A]](address A) netip.Addr {
	i, _ := netip.AddrFromSlice(address.Bytes())
	return i
}

func fromNetip(a netip.Addr) any {
	return ip.MustFromBytes(a.AsSlice()...)
}

func TestExampleParseUnknown(t *testing.T) {
	ExampleParseUnknown()
}

func ExampleParseUnknown() {
	examples := []string{"2001:db8::1", "192.0.2.1", "foobar"}
	for _, s := range examples {
		address, err := ip.ParseUnknown(s)
		if err != nil {
			println("Not address:", err.Error())
		} else {
			println("Address:", address.String())
		}
	}
}
