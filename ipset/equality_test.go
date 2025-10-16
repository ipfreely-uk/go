// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	. "github.com/ipfreely-uk/go/ipset"
	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	v4 := ip.V4()
	one := NewSingle(v4.FromInt(1))
	two := NewSingle(v4.FromInt(2))
	three := NewSingle(v4.FromInt(3))
	ten := NewSingle(v4.FromInt(10))
	eleven := NewSingle(v4.FromInt(11))
	oneTen := NewDiscrete(one, ten)
	oneEleven := NewDiscrete(one, eleven)
	{
		candidates := []Discrete[ip.Addr4]{
			one,
			NewDiscrete(one, two, three, ten),
			NewDiscrete(one, three, ten),
			NewDiscrete(one, two, ten),
			NewDiscrete(two, three, ten),
			oneTen,
			oneEleven,
		}

		for i, c0 := range candidates {
			for j, c1 := range candidates {
				expected := i == j
				actual := Eq(c0, c1)
				assert.Equal(t, expected, actual)
			}
		}
	}
	{
		set0 := NewDiscrete(two, three, ten)
		set1 := NewDiscrete(two, three, ten)
		actual := Eq(set0, set1)
		assert.True(t, actual)
	}
}
