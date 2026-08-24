module github.com/ipfreely-uk/go

go 1.26.0

// Test/example dependencies
require github.com/stretchr/testify v1.12.1

require go.yaml.in/yaml/v3 v3.0.5 // indirect

retract [v0.0.0-alpha, v0.0.37-beta] // old builds
