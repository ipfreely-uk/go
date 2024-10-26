/*
Network ranges and arbitrary IP address sets.

# Set Types

Interface hierarchy: [Block] set is [Interval] set is [Discrete] set.

[NewBlock] creates a RFC-4632 CIDR [Block] set. [NewSingle] creates a [Block] set from a single address.

[NewInterval] creates contiguous [Interval] set of addresses from lower and upper bounds.

[NewDiscrete] creates [Discrete] set as union of other address sets. Use [NewDiscrete] to create the empty set.

Functions return a struct that conforms to the most specialized interface possible.

# Iteration

Use [Discrete].Addresses to iterate over constituent addresses.

Use [Discrete].Intervals to iterate over constituent [Interval] sets.

Use the [Blocks] function to iterate over constituent [Block] sets within [Interval] sets.

# References

https://www.rfc-editor.org/rfc/rfc4632
*/
package ipset
