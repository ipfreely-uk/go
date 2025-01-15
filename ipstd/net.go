package ipstd

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import (
	"net"

	"github.com/ipfreely-uk/go/ip"
)

// Numeric type to `net` package
func ToIP(a ip.Address) net.IP {
	return a.Bytes()
}

// `net` package to numeric type
func FromIP(a net.IP) ip.Address {
	return ip.MustFromBytes(a...)
}

// Numeric type to `net` package
func ToIPMask(a ip.Address) net.IPMask {
	return a.Bytes()
}

// `net` package to numeric type
func FromIPMask(a net.IPMask) ip.Address {
	return ip.MustFromBytes(a...)
}
