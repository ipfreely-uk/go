package ip

// Tests equality
func Eq[A Int[A]](address0, address1 A) (areEqual bool) {
	return eq(address0, address1)
}

func eq(a, b any) bool {
	return a == b
}
