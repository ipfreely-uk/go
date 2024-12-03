package ipset

import "github.com/ipfreely-uk/go/ip"

func order[C ip.Int[C]](this, that C) (least, greatest C) {
	if this.Compare(that) <= 0 {
		return this, that
	}
	return that, this
}
