package ip_test

import (
	"github.com/ipfreely-uk/go/ip"
)

func ExampleV6() {
	bytes := make([]byte, 16)
	bytes[0] = 0xFE
	bytes[1] = 0x80
	address, _ := ip.V6().FromBytes(bytes...)
	println(address.String())
}
