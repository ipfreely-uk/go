ALL CODE IN ALPHA STATE

# IPFreely.uk

IP address manipulation library written in Go.
This library treats IP addresses as
[generic](https://go.dev/doc/tutorial/generics)
unsigned integers capable of arithmetic and bitwise operations.

## Packages

Add an import statement to [go.mod](https://go.dev/doc/modules/gomod-ref) to utilise.

| Package                                            | Purpose                                 |
|----------------------------------------------------|-----------------------------------------|
| `import github.com/ipfreely-uk/go/ip`              | Core IP address types                   |
| `import github.com/ipfreely-uk/go/ip/compare`      | Generic comparison types and functions  |
| `import github.com/ipfreely-uk/go/ip/network`      | IP address collection & iteration types |
| `import github.com/ipfreely-uk/go/ip/network/cidr` | CIDR notation functions                 |
| `import github.com/ipfreely-uk/go/ip/subnet`       | CIDR subnet functions                   |

## Versus Standard Library

Selective comparison with standard library types in
[net](https://pkg.go.dev/net@go1.22.2) and [netip](https://pkg.go.dev/net/netip@go1.22.2).

_TODO_

## Links

 - [IPFreely.uk Website](https://ipfreely.uk)
 - [Source Code](https://github.com/ipfreely-uk/go)
 - [Documentation](https://pkg.go.dev/github.com/ipfreely-uk/go)

## Continuous Integration

[![Go](https://github.com/ipfreely-uk/go/actions/workflows/go.yml/badge.svg)](https://github.com/ipfreely-uk/go/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/ipfreely-uk/go.svg)](https://pkg.go.dev/github.com/ipfreely-uk/go)
