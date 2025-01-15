package ipstd

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import (
	"net/netip"

	"github.com/ipfreely-uk/go/ip"
)

// Numeric type to `netip` package
func ToAddr(a ip.Address) netip.Addr {
	r, _ := netip.AddrFromSlice(a.Bytes())
	return r
}

// `netip` package to numeric type
func FromAddr(a netip.Addr) ip.Address {
	return ip.MustFromBytes(a.AsSlice()...)
}
