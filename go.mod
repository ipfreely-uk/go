module github.com/ipfreely-uk/go

go 1.23.1

// Test/example dependencies
require github.com/stretchr/testify v1.10.0

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract [v0.0.0-alpha, v0.0.34-beta] // old builds
