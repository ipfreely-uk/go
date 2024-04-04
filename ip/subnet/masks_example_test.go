package subnet_test

import (
	"fmt"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/subnet"
)

func Example_subnet_Mask() {
	network, _ := ip.V4().FromBytes(192, 168, 0, 0)
	mask := subnet.Mask(ip.V4(), 24)

	println("First: %s", network.String())
	println("Last: %s", mask.Not().Or(network).String())
	println("Mask: %s", mask.String())
}

func Example_subnet_AddressCount() {
	family4 := ip.V4()
	for mask := 0; mask <= family4.Width(); mask++ {
		count := subnet.AddressCount(family4, mask)
		msg := fmt.Sprintf("IPv%d /%d ==\t%s", family4.Version(), mask, count.String())
		println(msg)
	}
	family6 := ip.V6()
	for mask := 0; mask <= family6.Width(); mask++ {
		count := subnet.AddressCount(family6, mask)
		msg := fmt.Sprintf("IPv%d /%d ==\t%s", family6.Version(), mask, count.String())
		println(msg)
	}
}
