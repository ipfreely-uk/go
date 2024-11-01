package ipset

import "github.com/ipfreely-uk/go/ip"

func greatest[C ip.Int[C]](this, that C) (greatest C) {
	if this.Compare(that) >= 0 {
		return this
	}
	return that
}

func least[C ip.Int[C]](this, that C) (least C) {
	if this.Compare(that) <= 0 {
		return this
	}
	return that
}
