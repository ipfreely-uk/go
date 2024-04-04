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
	family := ip.V4()
	for mask := 0; mask <= family.Width(); mask++ {
		count := subnet.AddressCount(family, mask)
		msg := fmt.Sprintf("IPv%d /%d ==\t%s", family.Version(), mask, count.String())
		println(msg)
	}
	// family = ip.V6()
	// for mask := 0; mask <= family.Width(); mask++ {
	// 	count := subnet.AddressCount(family, mask)
	// 	msg := fmt.Sprintf("IPv%d /%d ==\t%s", family.Version(), mask, count.String())
	// 	println(msg)
	// }
}
