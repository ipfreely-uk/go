package ipset

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import (
	"iter"

	"github.com/ipfreely-uk/go/ip"
)

func emptySeq[E any]() iter.Seq[E] {
	return func(yield func(E) bool) {}
}

func singleSeq[E any](element E) iter.Seq[E] {
	return func(yield func(E) bool) {
		_ = yield(element)
	}
}

func sliceSeq[E any](slice []E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, e := range slice {
			more := yield(e)
			if !more {
				return
			}
		}
	}
}

func intervalToAddressSeq[A ip.Int[A]](slice []Interval[A]) iter.Seq[A] {
	return func(yield func(A) bool) {
		for _, i := range slice {
			for a := range i.Addresses() {
				more := yield(a)
				if !more {
					return
				}
			}
		}
	}
}
