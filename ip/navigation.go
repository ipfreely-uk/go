package ip

// Zero address for family.
func MinAddress[A Number[A]](fam Family[A]) (firstInFamily A) {
	return fam.FromInt(0)
}

// Maximum address for family
func MaxAddress[A Number[A]](fam Family[A]) (lastInFamily A) {
	return MinAddress(fam).Not()
}

// Increments argument by one with overflow
func Next[A Number[A]](address A) (incremented A) {
	one := address.Family().FromInt(1)
	return address.Add(one)
}

// Decrements argument by one with underflow
func Prev[A Number[A]](address A) (decremented A) {
	one := address.Family().FromInt(1)
	return address.Subtract(one)
}
