package ip_test

import (
	"github.com/ipfreely-uk/go/ip"
)

func Example_ip_FromBytes() {
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

func Example_ip_ParseUnknown() {
	examples := []string{"::1", "127.0.0.1", "foobar"}
	for _, s := range examples {
		address, err := ip.ParseUnknown(s)
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
}
