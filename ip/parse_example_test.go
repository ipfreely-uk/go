package ip_test

import (
	"github.com/ipfreely-uk/go/ip"
)

func ExampleFromBytes() {
	address, err := ip.FromBytes(127, 0, 0, 1)
	if err != nil {
		println("Not address: %v", err)
	}
	switch a := address.(type) {
	case ip.Address4:
		println("Not address: %v", a.String())
	case ip.Address6:
		println("Not address: %v", a.String())
	}
}

func ExampleParseUnknown() {
	examples := []string{"::1", "127.0.0.1", "foobar"}
	for _, s := range examples {
		address, err := ip.ParseUnknown(s)
		if err != nil {
			println("Not address: %s", err)
		}
		switch a := address.(type) {
		case ip.Address4:
			println("IPv4 address: %s", a.String())
		case ip.Address6:
			println("IPv6 address: %s", a.String())
		}
	}
}
