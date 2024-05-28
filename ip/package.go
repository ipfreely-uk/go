// IP addresses as generic, immutable positive integers.
// Use for arithmetic and bitwise operations.
//
// IP address [Family] types act as factories for [Address] values.
// Use [V4] and [V6] to obtain a [Family].
//
// Use [Family_FromBytes] or [FromBytes] to obtain [Address] from `[]byte`.
// Use [Parse] or [ParseUnknown] to obtain [Address] from `string`.
//
// Interfaces in this package cannot be implemented outside it.
package ip
