// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
)

func TestExampleV6(t *testing.T) {
	ExampleV6()
}

func ExampleV6() {
	family := ip.V6()
	bytes := make([]byte, family.Width()/8)
	bytes[0] = 0x20
	bytes[1] = 0x01
	bytes[2] = 0xDB
	bytes[3] = 0x80
	bytes[15] = 1

	address := family.MustFromBytes(bytes...)

	println(address.String())
}
