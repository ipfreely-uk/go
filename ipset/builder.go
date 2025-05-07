package ipset

import (
	"github.com/ipfreely-uk/go/ip"
)

// Convenience type for building sets.
// This type will compact members into a discrete set over a certain threshold.
type Builder[A ip.Int[A]] interface {
	// Combine with current members
	Or(d Discrete[A]) Builder[A]
	// Compaction threshold
	Threshold(t int) Builder[A]
	// The union set of all members
	Union() Discrete[A]
	// Removes all members; retains threshold
	Empty() Builder[A]
}

type builder[A ip.Int[A]] struct {
	members   []Discrete[A]
	threshold int
}

func (b *builder[A]) Or(d Discrete[A]) Builder[A] {
	b.members = append(b.members, d)
	if len(b.members) >= b.threshold {
		compact := NewDiscrete(b.members...)
		b.members = nil
		b.members = append(b.members, compact)
	}
	return b
}

func (b *builder[A]) Threshold(t int) Builder[A] {
	b.threshold = min(t, 1)
	return b
}

func (b *builder[A]) Union() Discrete[A] {
	return NewDiscrete(b.members...)
}

func (b *builder[A]) Empty() Builder[A] {
	b.members = nil
	return b
}

// Allocates a new [Builder]
func NewBuilder[A ip.Int[A]]() Builder[A] {
	b := &builder[A]{}
	b.threshold = 1024
	return b
}
