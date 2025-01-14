// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
)

func TestExampleNewInterval(t *testing.T) {
	ExampleNewInterval()
}

func ExampleNewInterval() {
	// 1st address in IPv4 subnet is network address.
	// Last address in IPv4 subnet is broadcast address.
	// The rest can be used for unicast.
	sub := ipset.NewBlock(ip.V4().MustFromBytes(203, 0, 113, 8), 29)
	first := ip.Next(sub.First())
	last := ip.Prev(sub.Last())

	allocations := ipset.NewInterval(first, last)

	println("Assignable addresses in ", sub.String())
	for a := range allocations.Addresses() {
		println(a.String())
	}
}
