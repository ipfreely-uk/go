package cidr_test

import (
	"testing"

	"github.com/ipfreely-uk/go/ip"
	"github.com/ipfreely-uk/go/ip/network"
	"github.com/ipfreely-uk/go/ip/network/cidr"
	"github.com/stretchr/testify/assert"
)

func TestNotation(t *testing.T) {
	a := ip.V6().FromInt(0)
	mask := ip.V6().Width()
	block := network.NewBlock(a, mask)
	actual := cidr.Notation(block)
	assert.Equal(t, "::/128", actual)
}
