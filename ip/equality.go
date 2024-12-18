package ip

// Tests equality
func Eq(address0, address1 Address) (areEqual bool) {
	return eq(address0, address1)
}

func eq(a, b any) bool {
	return a == b
}
