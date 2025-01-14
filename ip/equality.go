// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip

// Tests equality
func Eq(address0, address1 Address) (areEqual bool) {
	return eq(address0, address1)
}

func eq(a, b any) bool {
	return a == b
}
