package ip

// Zero address for family.
func MinAddress[A Address[A]](fam Family[A]) A {
	return fam.FromInt(0)
}

// Maximum address for family
func MaxAddress[A Address[A]](fam Family[A]) A {
	return MinAddress(fam).Not()
}

// Increments argument by one with overflow
func Next[A Address[A]](address A) A {
	one := address.Family().FromInt(1)
	return address.Add(one)
}

// Decrements argument by one with underflow
func Prev[A Address[A]](address A) A {
	one := address.Family().FromInt(1)
	return address.Subtract(one)
}
