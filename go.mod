module github.com/ipfreely-uk/go

go 1.26.0

// Test/example dependencies
require github.com/stretchr/testify v1.12.0

require gopkg.in/yaml.v3 v3.0.1 // indirect

retract [v0.0.0-alpha, v0.0.37-beta] // old builds
