// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ip_test

import (
	"testing"

	. "github.com/ipfreely-uk/go/ip"
	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	{
		v4 := V4()
		one := v4.FromInt(1)
		two := v4.FromInt(2)
		assert.True(t, Eq(one, v4.FromInt(1)))
		assert.True(t, Eq(one, one))
		assert.False(t, Eq(one, two))
		assert.False(t, Eq(two, one))
	}
	{
		v6 := V6()
		one := v6.FromInt(1)
		complement := one.Not()
		complement2 := one.Not()
		assert.True(t, Eq(one, v6.FromInt(1)))
		assert.True(t, Eq(one, one))
		assert.True(t, Eq(complement, complement))
		assert.True(t, Eq(complement, complement2))
		assert.False(t, Eq(one, complement))
		assert.False(t, Eq(complement, one))
	}
}
