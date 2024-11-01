package ip

// Tests equality
func Eq[A Int[A]](address0, address1 A) (areEqual bool) {
	return address0.Compare(address1) == 0
}
