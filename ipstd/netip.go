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
