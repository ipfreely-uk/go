package ipset

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import "github.com/ipfreely-uk/go/ip"

func order[C ip.Int[C]](this, that C) (least, greatest C) {
	if this.Compare(that) <= 0 {
		return this, that
	}
	return that, this
}
