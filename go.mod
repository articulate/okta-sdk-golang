module github.com/articulate/okta-sdk-golang

go 1.12

require (
	github.com/go-yaml/yaml v2.1.0+incompatible
	github.com/kelseyhightower/envconfig v1.3.0
	github.com/okta/okta-sdk-golang v0.1.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/stretchr/testify v1.2.2
)

// For tests to run properly
replace github.com/okta/okta-sdk-golang => ./
