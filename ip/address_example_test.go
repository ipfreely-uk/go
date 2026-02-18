// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
)

func TestExampleNumber(t *testing.T) {
	ExampleInt()
}

func ExampleInt() {
	printN(ip.V4().MustFromBytes(192, 0, 2, 0), 3)
	printN(ip.MustParse(ip.V6(), "2001:db8::"), 3)
}

func printN[A ip.Int[A]](address A, n int) {
	a := address
	one := a.Family().FromInt(1)
	for range n {
		println(a.String())
		a = a.Add(one)
	}
}

func TestExampleAddress(t *testing.T) {
	ExampleAddress()
}

func ExampleAddress() {
	examples := []string{
		"2001:db8::",
		"192.0.2.0",
	}

	for _, e := range examples {
		address := ip.MustParseUnknown(e)
		switch a := address.(type) {
		case ip.Addr4:
			printNthAfter(a, 255)
		case ip.Addr6:
			printNthAfter(a, 0xFFFFFFFF)
		}
	}
}

func printNthAfter[A ip.Int[A]](address A, n uint32) {
	operand := address.Family().FromInt(n)
	result := address.Add(operand)
	println(result.String())
}
