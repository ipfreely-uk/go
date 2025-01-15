package ip

// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0

// Basic specification information.
type Spec interface {
	// Structs that conform to this interface must be produced by this package
	sealed()
	// IP address version
	Version() (constant Version)
	// Address width in bits - 32 or 128
	Width() (bits int)
	// Informational
	String() string
}

// IP address family.
// Obtain via [V4] or [V6] functions.
type Family[A Address] interface {
	Spec
	// Create address from bytes.
	// Returns error if slice is not [Width]/8 bytes.
	FromBytes(...byte) (address A, err error)
	// As FromBytes but panics on error
	MustFromBytes(...byte) (address A)
	// Create address from unsigned integer.
	// All values are valid.
	// For V4 operand 1 returns "0.0.0.1".
	// For V6 operand 1 return "::1".
	FromInt(i uint32) (address A)
}
