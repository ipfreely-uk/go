// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipstd

import (
	"net/netip"

	"github.com/ipfreely-uk/go/ip"
)

func ToAddr(a ip.Address) netip.Addr {
	r, _ := netip.AddrFromSlice(a.Bytes())
	return r
}

func FromAddr(a netip.Addr) ip.Address {
	return ip.MustFromBytes(a.AsSlice()...)
}
