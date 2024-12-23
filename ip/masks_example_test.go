package ip_test

import (
	"fmt"
	"testing"

	"github.com/dustin/go-humanize"
	"github.com/ipfreely-uk/go/ip"
)

func TestExampleSubnetMask(t *testing.T) {
	ExampleSubnetMask()
	ExampleSubnetMask_second()
}

func ExampleSubnetMask() {
	network := ip.V4().MustFromBytes(192, 0, 2, 128)
	printNetworkDetails(network, 26)
}

func printNetworkDetails[A ip.Int[A]](network A, maskBits int) {
	fam := network.Family()
	mask := ip.SubnetMask(fam, maskBits)
	inverseMask := mask.Not()

	zero := fam.FromInt(0)
	if !ip.Eq(mask.And(inverseMask), zero) {
		panic("Mask does not cover network address")
	}

	println("First Address:", network.String())
	println("Last Address:", network.Or(inverseMask).String())
	println("Mask:", mask.String())
	println(fmt.Sprintf("CIDR Notation: %s/%d", network.String(), maskBits))
}

func ExampleSubnetMask_second() {
	printAllMasks(ip.V4())
	printAllMasks(ip.V6())
}

func printAllMasks[A ip.Int[A]](f ip.Family[A]) {
	println(f.String())
	for bits := 0; bits <= f.Width(); bits++ {
		mask := ip.SubnetMask(f, bits)
		cidrTail := fmt.Sprintf("/%d", bits)
		println(mask.String(), "==", cidrTail)
	}
}

func TestExampleSubnetAddressCount(t *testing.T) {
	ExampleSubnetAddressCount()
}

func ExampleSubnetAddressCount() {
	printSubnetSizesForMasks(ip.V4())
	printSubnetSizesForMasks(ip.V6())
}

func printSubnetSizesForMasks[A ip.Int[A]](f ip.Family[A]) {
	for mask := 0; mask <= f.Width(); mask++ {
		count := ip.SubnetAddressCount(f, mask)
		msg := fmt.Sprintf("IPv%d /%d == %s", f.Version(), mask, humanize.BigComma(count))
		println(msg)
	}
}

func TestExampleSubnetMaskSize(t *testing.T) {
	ExampleSubnetMaskSize()
	ExampleSubnetMaskSize_second()
}

func ExampleSubnetMaskSize() {
	family := ip.V4()
	first := family.MustFromBytes(192, 0, 2, 0)
	last := family.MustFromBytes(192, 0, 2, 255)

	maskBits := ip.SubnetMaskSize(first, last)
	mask := ip.SubnetMask(family, maskBits)

	println(fmt.Sprintf("/%d", maskBits), "==", mask.String())
}

func ExampleSubnetMaskSize_second() {
	family := ip.V4()
	first := family.MustFromBytes(192, 0, 2, 0)
	last := family.MustFromBytes(192, 0, 2, 255)

	maskBits := ip.SubnetMaskSize(first, last)
	if maskBits != -1 {
		cidrNotation := fmt.Sprintf("%s/%d", first.String(), maskBits)
		println(first.String(), "-", last.String(), " is valid subnet ", cidrNotation)
	} else {
		println(first.String(), "-", last.String(), " is not a valid subnet")
	}
}

func TestExampleSubnetMaskCovers(t *testing.T) {
	ExampleSubnetMaskCovers()
}

func ExampleSubnetMaskCovers() {
	v4 := ip.V4()
	netAddress := v4.MustFromBytes(192, 0, 2, 0)

	for mask := 32; mask >= 20; mask-- {
		addrStr := netAddress.String()
		cidrNotation := fmt.Sprintf("%s/%d", addrStr, mask)
		if ip.SubnetMaskCovers(mask, netAddress) {
			println(cidrNotation, "is a valid expression")
		} else {
			maskAddr := ip.SubnetMask(v4, mask).String()
			println(cidrNotation, "is not a valid expression;", maskAddr, "does not cover", addrStr)
		}
	}
}

func TestExampleIsSubnetMask(t *testing.T) {
	ExampleIsSubnetMask()
}

func ExampleIsSubnetMask() {
	mask := ip.MustParse(ip.V4(), "255.255.255.0")
	bits := calculateSubnetBits(mask)
	msg := fmt.Sprintf("Subnets with mask %s are /%d networks", mask.String(), bits)
	println(msg)
}

// Calculate the /n CIDR expression for a given mask
func calculateSubnetBits[A ip.Int[A]](mask A) (bits int) {
	if !ip.IsSubnetMask(mask) {
		msg := fmt.Sprintf("%s is not a subnet mask", mask.String())
		panic(msg)
	}
	f := mask.Family()
	return f.Width() - mask.TrailingZeros()
}
