package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
)

func TestExampleV4(t *testing.T) {
	ExampleV4()
}

func ExampleV4() {
	address := ip.V4().MustFromBytes(203, 0, 113, 1)
	println(address.String())
}
