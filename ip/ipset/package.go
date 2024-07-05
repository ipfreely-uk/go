// Network ranges and IP address sets.
//
// [NewInterval] creates contiguous [Interval] set.
// [NewBlock] creates RFC-4632 https://www.rfc-editor.org/rfc/rfc4632 CIDR [Block] set.
// [NewDiscrete] creates non-contigous [Discrete].
//
// Use [Discrete.Addresses] to iterate constituent addresses.
// Use [Discrete.Ranges] to iterate constituent [Interval] types.
// Use [Blocks] to iterate constituent [Block] types.
package ipset
