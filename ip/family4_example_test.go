package ip_test

import (
	"github.com/ipfreely-uk/go/ip"
)

func ExampleV4() {
	address := ip.V4().MustFromBytes(203, 0, 113, 1)
	println(address.String())
}
