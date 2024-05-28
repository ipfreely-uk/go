// Network ranges and IP address sets.
//
// [NewRange] creates contiguous [AddressRange] set.
// [NewBlock] creates [RFC-4632](https://www.rfc-editor.org/rfc/rfc4632) CIDR [Block] set.
// [NewSet] creates non-contigous [AddressSet].
//
// Use [AddressSet_Addresses] to iterate over constituent addresses.
// Use [AddressSet_Ranges] to iterate over constituent [AddressRange]s.
// Use [Blocks] to iterate over constituent [Block]s.
package network
