package ipstd

import (
	"net"

	"github.com/ipfreely-uk/go/ip"
)

func ToIP(a ip.Address) net.IP {
	return a.Bytes()
}

func FromIP(a net.IP) ip.Address {
	return ip.MustFromBytes(a...)
}

func ToIPMask(a ip.Address) net.IPMask {
	return a.Bytes()
}

func FromIPMask(a net.IPMask) ip.Address {
	return ip.MustFromBytes(a...)
}
