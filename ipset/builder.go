package ipset

import (
	"slices"

	"github.com/ipfreely-uk/go/ip"
)

const threshold = 512

// [Discrete] set builder.
//
// The only advantage over using a slice and [NewDiscrete]
// is that this type will attempt to reduce the buffer over a certain threshold.
type Builder[A ip.Int[A]] struct {
	members []Discrete[A]
}

// Creates union with existing members
func (b *Builder[A]) Union(d Discrete[A]) *Builder[A] {
	b.members = append(b.members, d)
	b.compact()
	return b
}

func (b *Builder[A]) compact() {
	if len(b.members) >= threshold {
		compacted := NewDiscrete(b.members...)
		b.members = slices.Delete(b.members, 0, len(b.members))
		b.members = append(b.members, compacted)
	}
}

// Creates [Discrete] set and empties the buffer for reuse
func (b *Builder[A]) Build() Discrete[A] {
	d := NewDiscrete(b.members...)
	b.members = slices.Delete(b.members, 0, len(b.members))
	return d
}
