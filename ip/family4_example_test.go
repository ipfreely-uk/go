package ip_test

import (
	"github.com/ipfreely-uk/go/ip"
)

func Example_ip_V4() {
	address, _ := ip.V4().FromBytes(192, 168, 0, 1)
	println(address.String())
}
