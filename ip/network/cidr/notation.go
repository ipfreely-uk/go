package cidr

import (
	"fmt"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
)

func Notation[A ip.Address[A]](b network.Block[A]) string {
	return fmt.Sprintf("%s/%d", b.First(), b.MaskSize())
}

func Parse[A ip.Address[A]](f ip.Family[A], network string) (network.Block[A], error) {
	panic("TODO")
}
