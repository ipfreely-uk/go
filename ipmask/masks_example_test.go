// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipmask_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipmask"
	"github.com/ipfreely-uk/go/txt"
)

func TestExampleFor(t *testing.T) {
	ExampleFor()
	ExampleFor_second()
}

func ExampleFor() {
	network4 := ip.V4().MustFromBytes(192, 0, 2, 128)
	printNetworkDetails(network4, 26)

	println()

	network6 := ip.MustParse(ip.V6(), "2001:DB80::")
	printNetworkDetails(network6, 65)
}

func printNetworkDetails[A ip.Int[A]](network A, maskBits int) {
	fam := network.Family()
	mask := ipmask.For(fam, maskBits)
	maskComplement := mask.Not()

	zero := fam.FromInt(0)
	if !ip.Eq(mask.And(maskComplement), zero) {
		panic("Mask does not cover network address")
	}

	println("First Address:", network.String())
	println("Last Address:", network.Or(maskComplement).String())
	println("Mask:", mask.String())
	fmt.Printf("CIDR Notation: %s/%d\n", network.String(), maskBits)
}

func ExampleFor_second() {
	printAllMasks(ip.V4())
	printAllMasks(ip.V6())
}

func printAllMasks[A ip.Int[A]](f ip.Family[A]) {
	println(f.String())
	for bits := 0; bits <= f.Width(); bits++ {
		mask := ipmask.For(f, bits)
		cidrTail := fmt.Sprintf("/%d", bits)
		println(mask.String(), "==", cidrTail)
	}
}

func TestExampleSize(t *testing.T) {
	ExampleSize()
	ExampleSize_second()
}

func ExampleSize() {
	printSubnetSizesForMasks(ip.V4())
	printSubnetSizesForMasks(ip.V6())
}

func printSubnetSizesForMasks[A ip.Int[A]](f ip.Family[A]) {
	for mask := 0; mask <= f.Width(); mask++ {
		count := ipmask.Size(f, mask)
		fmt.Printf("IPv%d /%d = %s\n", f.Version(), mask, txt.CommaDelim(count))
	}
}

func ExampleSize_second() {
	family := ip.V4()
	min := big.NewInt(50)
	bits := minimumMaskThatSatisfies(family, min)
	mask := ipmask.For(family, bits).String()
	count := txt.CommaDelim(min)
	msg := fmt.Sprintf("/%d network (%s) is the minimum size that can allocate %s addresses", bits, mask, count)
	println(msg)
}

func minimumMaskThatSatisfies[A ip.Int[A]](f ip.Family[A], allocatable *big.Int) int {
	var min *big.Int
	// IPv4 reserves 1st & last
	if f.Version() == ip.Version4 {
		two := big.NewInt(2)
		min = two.Add(two, allocatable)
	} else {
		min = allocatable
	}
	for i := f.Width(); i >= 0; i-- {
		s := ipmask.Size(f, i)
		if min.Cmp(s) <= 0 {
			return i
		}
	}
	panic("illegal argument")
}

func TestExampleTest(t *testing.T) {
	ExampleTest()
	ExampleTest_second()
}

func ExampleTest() {
	family := ip.V4()
	first := family.MustFromBytes(192, 0, 2, 0)
	last := family.MustFromBytes(192, 0, 2, 255)

	maskBits := ipmask.Test(first, last)
	mask := ipmask.For(family, maskBits)

	println(fmt.Sprintf("/%d", maskBits), "==", mask.String())
}

func ExampleTest_second() {
	family := ip.V4()
	first := family.MustFromBytes(192, 0, 2, 0)
	last := family.MustFromBytes(192, 0, 2, 255)

	maskBits := ipmask.Test(first, last)
	if maskBits != -1 {
		cidrNotation := fmt.Sprintf("%s/%d", first.String(), maskBits)
		println(first.String(), "-", last.String(), " is valid subnet ", cidrNotation)
	} else {
		println(first.String(), "-", last.String(), " is not a valid subnet")
	}
}

func TestExampleCovers(t *testing.T) {
	ExampleCovers()
}

func ExampleCovers() {
	v4 := ip.V4()
	netAddress := v4.MustFromBytes(192, 0, 2, 0)

	for mask := 32; mask >= 20; mask-- {
		addrStr := netAddress.String()
		cidrNotation := fmt.Sprintf("%s/%d", addrStr, mask)
		if ipmask.Covers(mask, netAddress) {
			println(cidrNotation, "is a valid expression")
		} else {
			maskAddr := ipmask.For(v4, mask).String()
			println(cidrNotation, "is not a valid expression;", maskAddr, "does not cover", addrStr)
		}
	}
}

func TestExampleBits(t *testing.T) {
	ExampleBits()
}

func ExampleBits() {
	v4 := ip.V4()
	v6 := ip.V6()
	validateMask(v4.MustFromBytes(0, 0, 0, 0))
	validateMask(v4.MustFromBytes(192, 168, 0, 0))
	validateMask(v4.MustFromBytes(255, 255, 255, 0))
	validateMask(v4.MustFromBytes(255, 255, 0xF0, 0))
	validateMask(v4.MustFromBytes(255, 255, 0x80, 0))
	validateMask(v4.MustFromBytes(255, 255, 0xF, 0))
	validateMask(ipmask.For(v6, 56))
}

func validateMask[A ip.Int[A]](address A) {
	bits := ipmask.Bits(address)
	ver := address.Version()
	if bits >= 0 {
		msg := fmt.Sprintf("%s is valid IPv%d mask /%d", address.String(), ver, bits)
		println(msg)
	} else {
		msg := fmt.Sprintf("%s is not a valid mask", address.String())
		println(msg)
	}
}
