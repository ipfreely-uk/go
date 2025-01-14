// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
)

func TestExampleV4(t *testing.T) {
	ExampleV4()
}

func ExampleV4() {
	address := ip.V4().MustFromBytes(203, 0, 113, 1)
	println(address.String())
}
