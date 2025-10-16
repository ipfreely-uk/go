// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipstd_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	. "github.com/ipfreely-uk/go/ipstd"
	"github.com/stretchr/testify/assert"
)

func TestToAddr(t *testing.T) {
	expected := ip.MustParse(ip.V4(), "127.0.0.1")
	n := ToAddr(expected)
	actual := FromAddr(n)
	assert.Equal(t, expected, actual)
}
