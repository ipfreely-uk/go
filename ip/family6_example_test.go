package ip_test

import (
	"github.com/ipfreely-uk/go/ip"
)

func ExampleV6() {
	family := ip.V6()
	bytes := make([]byte, family.Width()/8)
	bytes[0] = 0x20
	bytes[1] = 0x01
	bytes[2] = 0xDB
	bytes[3] = 0x80
	bytes[15] = 1
	address, _ := family.FromBytes(bytes...)
	println(address.String())
}
