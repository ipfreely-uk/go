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
	nip := netip.MustParseAddr("2001:db8::")

	address := ip.MustFromBytes(nip.AsSlice()...)

	switch a := address.(type) {
	case ip.Addr4:
		println("IPv4 address:", a.String())
	case ip.Addr6:
		println("IPv6 address", a.String())
	}
}

func TestExampleParseUnknown(t *testing.T) {
	ExampleParseUnknown()
}

func ExampleParseUnknown() {
	examples := []string{"2001:db8::1", "192.0.2.1", "foobar"}
	for _, s := range examples {
		address, err := ip.ParseUnknown(s)
		if err != nil {
			println("Not address:", err)
		}
		switch a := address.(type) {
		case ip.Addr4:
			println("IPv4 address:", a.String())
		case ip.Addr6:
			println("IPv6 address:", a.String())
		}
	}
}
