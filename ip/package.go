/*
IP addresses as generic, immutable positive integers.
Use for arithmetic and bitwise operations.

# IP Address Types

IP address [Family] types act as factories for the [Int] structs [Addr4] and [Addr6].
Use [V4] or [V6] to obtain a [Family].

Use [Family].FromBytes or [FromBytes] to obtain [Int] from byte slice.
Use [Parse] or [ParseUnknown] to obtain [Int] from string.
*/
package ip
