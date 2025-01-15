package ipset

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

import (
	"iter"
	"math/big"

	"github.com/ipfreely-uk/go/ip"
)

type empty[A ip.Int[A]] struct{}

func (e *empty[A]) Contains(A) bool {
	return false
}

func (e *empty[A]) Size() *big.Int {
	return big.NewInt(0)
}

func (e *empty[A]) Addresses() iter.Seq[A] {
	return emptySeq[A]()
}

func (e *empty[A]) Intervals() iter.Seq[Interval[A]] {
	return emptySeq[Interval[A]]()
}

func (e *empty[A]) String() string {
	return "{}"
}

func emptySet[A ip.Int[A]]() Discrete[A] {
	return &empty[A]{}
}
