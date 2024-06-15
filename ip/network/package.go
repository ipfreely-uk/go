// Network ranges and IP address sets.
//
// [NewRange] creates contiguous [AddressRange] set.
// [NewBlock] creates RFC-4632 https://www.rfc-editor.org/rfc/rfc4632 CIDR [Block] set.
// [NewSet] creates non-contigous [AddressSet].
//
// Use [AddressSet.Addresses] to iterate constituent addresses.
// Use [AddressSet.Ranges] to iterate constituent [AddressRange] types.
// Use [Blocks] to iterate constituent [Block] types.
package network
