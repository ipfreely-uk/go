package subnet_test

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/subnet"
)

func ExampleMask() {
	network := ip.V4().MustFromBytes(192, 168, 0, 0)

	maskBits := 24
	mask := subnet.Mask(ip.V4(), maskBits)

	println("First:", network.String())
	println("Last:", mask.Not().Or(network).String())
	println("Mask:", mask.String())
}

func ExampleMask_second() {
	printAllMasks(ip.V4())
	printAllMasks(ip.V6())
}

func printAllMasks[A ip.Address[A]](f ip.Family[A]) {
	println(fmt.Sprintf("IPv%d", f.Version()))
	for bits := 0; bits <= f.Width(); bits++ {
		mask := subnet.Mask(f, bits)
		cidrTail := fmt.Sprintf("/%d", bits)
		println(mask.String(), "==", cidrTail)
	}
}

func ExampleAddressCount() {
	printSubnetSizesForMasks(ip.V4())
	printSubnetSizesForMasks(ip.V6())
}

func printSubnetSizesForMasks[A ip.Address[A]](family ip.Family[A]) {
	for mask := 0; mask <= family.Width(); mask++ {
		count := subnet.AddressCount(family, mask)
		msg := fmt.Sprintf("IPv%d /%d == %s", family.Version(), mask, humanize.BigComma(count))
		println(msg)
	}
}

func ExampleMaskSize() {
	family := ip.V4()
	first := family.MustFromBytes(192, 0, 2, 0)
	last := family.MustFromBytes(192, 0, 2, 255)

	maskBits := subnet.MaskSize(first, last)
	mask := subnet.Mask(family, maskBits)

	println(fmt.Sprintf("/%d", maskBits), "==", mask.String())
}

func ExampleMaskSize_second() {
	family := ip.V4()
	first := family.MustFromBytes(192, 0, 2, 0)
	last := family.MustFromBytes(192, 0, 2, 255)

	maskBits := subnet.MaskSize(first, last)
	if maskBits != -1 {
		cidrNotation := fmt.Sprintf("%s/%d", first.String(), maskBits)
		println(first.String(), "-", last.String(), " is valid subnet ", cidrNotation)
	} else {
		println(first.String(), "-", last.String(), " is not a valid subnet")
	}
}
