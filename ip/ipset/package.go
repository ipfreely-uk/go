/*
Network ranges and arbitrary IP address sets.

# Set Types

[NewInterval] creates contiguous [Interval] set of addresses from lower and upper bounds.
[NewBlock] creates a RFC-4632 CIDR [Block] set.
[NewSingle] creates a [Block] set from a single address.
[NewDiscrete] creates [Discrete] set of addresses from arbitrary sets.
Use [NewDiscrete] to create the empty set.

# Iteration

Use [Discrete].Addresses to iterate over constituent addresses.
Use [Discrete].Intervals to iterate over constituent [Interval] sets.
Use the [Blocks] function to iterate over constituent [Block] sets within [Interval] sets.

# References

https://www.rfc-editor.org/rfc/rfc4632
*/
package ipset
