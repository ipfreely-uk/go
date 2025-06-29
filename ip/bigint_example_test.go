// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/txt"
)

func TestExampleToBigInt(t *testing.T) {
	ExampleToBigInt()
}

func ExampleToBigInt() {
	printNumberOfAddresses(ip.MinAddress(ip.V4()), ip.MaxAddress(ip.V4()))
	printNumberOfAddresses(ip.MinAddress(ip.V6()), ip.MaxAddress(ip.V6()))
}

func printNumberOfAddresses[A ip.Int[A]](first, last A) {
	diff := last.Subtract(first)

	n := ip.ToBigInt(diff)
	rangeSize := n.Add(n, big.NewInt(1))

	fmt.Printf("%v to %v = %s addresses\n", first, last, txt.CommaDelim(rangeSize))
}
