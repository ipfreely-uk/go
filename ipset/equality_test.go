// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipset_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ipset"
	"github.com/stretchr/testify/assert"
)

func TestEq(t *testing.T) {
	v4 := ip.V4()
	one := ipset.NewSingle(v4.FromInt(1))
	two := ipset.NewSingle(v4.FromInt(2))
	three := ipset.NewSingle(v4.FromInt(3))
	ten := ipset.NewSingle(v4.FromInt(10))
	eleven := ipset.NewSingle(v4.FromInt(11))
	oneTen := ipset.NewDiscrete(one, ten)
	oneEleven := ipset.NewDiscrete(one, eleven)
	{
		candidates := []ipset.Discrete[ip.Addr4]{
			one,
			ipset.NewDiscrete(one, two, three, ten),
			ipset.NewDiscrete(one, three, ten),
			ipset.NewDiscrete(one, two, ten),
			ipset.NewDiscrete(two, three, ten),
			oneTen,
			oneEleven,
		}

		for i, c0 := range candidates {
			for j, c1 := range candidates {
				expected := i == j
				actual := ipset.Eq(c0, c1)
				assert.Equal(t, expected, actual)
			}
		}
	}
	{
		set0 := ipset.NewDiscrete(two, three, ten)
		set1 := ipset.NewDiscrete(two, three, ten)
		actual := ipset.Eq(set0, set1)
		assert.True(t, actual)
	}
}
