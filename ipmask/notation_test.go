// Copyright 2024-2025 https://github.com/ipfreely-uk/go/blob/main/LICENSE
// SPDX-License-Identifier: Apache-2.0
package ipmask_test

import (
	"fmt"
	"testing"

	"github.com/ipfreely-uk/go/ip"
	. "github.com/ipfreely-uk/go/ipmask"
	"github.com/stretchr/testify/assert"
)

func TestParseCIDRNotation(t *testing.T) {
	{
		legal := []string{
			"192.168.0.0/24",
			"192.168.0.0/32",
		}
		for _, c := range legal {
			a, m, err := ParseCIDRNotation(ip.V4(), c)
			assert.Nil(t, err)
			cidr := fmt.Sprintf("%s/%d", a, m)
			assert.Equal(t, c, cidr)
		}
	}
	{
		legal := []string{
			"::/0",
			"::/128",
		}
		for _, c := range legal {
			a, m, err := ParseCIDRNotation(ip.V6(), c)
			assert.Nil(t, err)
			cidr := fmt.Sprintf("%s/%d", a, m)
			assert.Equal(t, c, cidr)
		}
	}
	{
		illegal := []string{
			"foo/24",
			"192.168.0.0/128",
			"192.168.0.0/0",
			"192.168.0.0/",
			"192.168.0.0",
			"::/-1",
			"::/129",
		}
		for _, c := range illegal {
			_, _, err := ParseCIDRNotation(ip.V4(), c)
			assert.NotNil(t, err)
		}
		for _, c := range illegal {
			_, _, err := ParseCIDRNotation(ip.V6(), c)
			assert.NotNil(t, err)
		}
	}
}

func TestParseUnknownCIDRNotation(t *testing.T) {
	legal := []string{
		"192.168.0.0/24",
		"192.168.0.0/32",
		"::/0",
		"::/128",
	}
	for _, c := range legal {
		a, m, err := ParseUnknownCIDRNotation(c)
		assert.Nil(t, err)
		cidr := fmt.Sprintf("%s/%d", a.String(), m)
		assert.Equal(t, c, cidr)
	}
	illegal := []string{
		"foo/24",
		"192.168.0.0/128",
		"192.168.0.0/0",
		"192.168.0.0/",
		"192.168.0.0",
		"::/-1",
		"::/129",
	}
	for _, c := range illegal {
		_, _, err := ParseUnknownCIDRNotation(c)
		assert.NotNil(t, err)
	}
}
