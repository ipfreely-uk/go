# IPFreely.uk

IP address manipulation library written in Go.
This library treats IP addresses as
[generic](https://go.dev/doc/tutorial/generics)
unsigned integers capable of arithmetic and bitwise operations
and includes a few discrete set collection types.
This library does not perform network I/O.

## Example

```go
package main

import(
    "fmt"

    "github.com/ipfreely-up/go/ip"
)

func main() {
	network4 := ip.V4().MustFromBytes(192, 0, 2, 128)
	printNetworkDetails(network4, 26)

	println()

	network6 := ip.MustParse(ip.V6(), "2001:DB80::")
	printNetworkDetails(network6, 65)
}

func printNetworkDetails[A ip.Int[A]](network A, maskBits int) {
	fam := network.Family()
	mask := ip.SubnetMask(fam, maskBits)
	maskComplement := mask.Not()

	zero := fam.FromInt(0)
	if !ip.Eq(mask.And(maskComplement), zero) {
		panic("Mask does not cover network address")
	}

	println("First Address:", network.String())
	println("Last Address:", network.Or(maskComplement).String())
	println("Mask:", mask.String())
	fmt.Printf("CIDR Notation: %s/%d\n", network.String(), maskBits)
}
```

Output:

```
First Address: 192.0.2.128
Last Address: 192.0.2.191
Mask: 255.255.255.192
CIDR Notation: 192.0.2.128/26

First Address: 2001:db80::
Last Address: 2001:db80::7fff:ffff:ffff:ffff
Mask: ffff:ffff:ffff:ffff:8000::
CIDR Notation: 2001:db80::/65
```

## Packages

Add an import statement to [go.mod](https://go.dev/doc/modules/gomod-ref) to utilise.

| Package                                   | Purpose                  |
|-------------------------------------------|--------------------------|
| `import github.com/ipfreely-uk/go/ip`     | IP addresses as integers |
| `import github.com/ipfreely-uk/go/ipmask` | IP subnet mask funcs     |
| `import github.com/ipfreely-uk/go/ipset`  | IP address discrete sets |
| `import github.com/ipfreely-uk/go/ipstd`  | Convert to/from stdlib   |

## Versus Standard Library

Selective comparison with standard library types in [netip](https://pkg.go.dev/net/netip).

| Feature                | IPFreely.uk | netip |
| -----------------------|-------------|-------|
| Immutable Types        | Y           | Y     |
| Categorization         |             | Y     |
| Generic Types          | Y           |       |
| Arithmetic/Bitwise Ops | Y           |       |
| IPv6 Zones             |             | Y     |
| CIDR Blocks            | Y           | Y     |
| Arbitrary Ranges/Sets  | Y           |       |
| Iteration              | Y           |       |

## Links

 - [IPFreely.uk Website](https://ipfreely.uk)
 - [Source Code](https://github.com/ipfreely-uk/go)
 - [Documentation](https://pkg.go.dev/github.com/ipfreely-uk/go)

## Continuous Integration

[![Go](https://github.com/ipfreely-uk/go/actions/workflows/go.yml/badge.svg)](https://github.com/ipfreely-uk/go/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/ipfreely-uk/go.svg)](https://pkg.go.dev/github.com/ipfreely-uk/go)
