package ip_test

import (
	"github.com/ipfreely-uk/go/ip"
)

func ExampleV4() {
	address, _ := ip.V4().FromBytes(203, 0, 113, 1)
	println(address.String())
}
