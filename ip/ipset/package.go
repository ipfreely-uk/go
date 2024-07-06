// Network ranges and arbitrary IP address sets.
//
// [NewInterval] creates contiguous [Interval] set of addresses.
// [NewBlock] creates RFC-4632 https://www.rfc-editor.org/rfc/rfc4632 CIDR [Block] set.
// [NewDiscrete] creates non-contigous [Discrete] set of addresses.
//
// Use [Discrete.Addresses] to iterate over constituent addresses.
// Use [Discrete.Intervals] to iterate over constituent [Interval] types.
// Use [Blocks] to iterate over constituent [Block] types within an [Interval] type.
package ipset
