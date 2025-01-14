// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset_test

import (
	"fmt"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
)

func TestExampleNewDiscrete(t *testing.T) {
	ExampleNewDiscrete()
}

func ExampleNewDiscrete() {
	set := func(first, last string) ipset.Interval[ip.Addr4] {
		v4 := ip.V4()
		p := ip.MustParse[ip.Addr4]
		return ipset.NewInterval(p(v4, first), p(v4, last))
	}
	r0 := set("192.0.2.0", "192.0.2.100")
	r1 := set("192.0.2.101", "192.0.2.111")
	r2 := set("192.0.2.200", "192.0.2.200")

	union := ipset.NewDiscrete(r0, r1, r2)
	println(r0.String(), "\u222A", r1.String(), "\u222A", r2.String(), "=", union.String())
}

func TestExampleNewDiscrete_second(t *testing.T) {
	ExampleNewDiscrete_second()
}

func ExampleNewDiscrete_second() {
	printEmptySetFor(ip.V4())
	printEmptySetFor(ip.V6())
}

func printEmptySetFor[A ip.Int[A]](f ip.Family[A]) {
	empty := ipset.NewDiscrete[A]()
	println(f.String(), empty.String())
}

func TestExampleNewDiscrete_third(t *testing.T) {
	ExampleNewDiscrete_third()
}

func ExampleNewDiscrete_third() {
	s0 := parseV4("192.0.2.0/32")
	s1 := parseV4("192.0.2.11/32")
	s2 := parseV4("192.0.2.12/32")
	printSetType(s0)
	printSetType(s1, s2)
	printSetType(s0, s1, s2)
}

func parseV4(notation string) ipset.Block[ip.Addr4] {
	a, m, err := ipset.ParseCIDRNotation(ip.V4(), notation)
	if err != nil {
		panic(err)
	}
	return ipset.NewBlock(a, m)
}

func printSetType[A ip.Int[A]](sets ...ipset.Discrete[A]) {
	union := ipset.NewDiscrete(sets...)
	switch s := union.(type) {
	case ipset.Block[A]:
		println(fmt.Sprintf("%s is a block set", s.String()))
	case ipset.Interval[A]:
		println(fmt.Sprintf("%s is an interval set", s.String()))
	case ipset.Discrete[A]:
		println(fmt.Sprintf("%s is a discrete set", s.String()))
	}
}
