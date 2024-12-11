# Toggl API
[![Go Reference](https://pkg.go.dev/badge/github.com/go-api-libs/toggl.svg)](https://pkg.go.dev/github.com/go-api-libs/toggl/pkg/toggl)
[![OpenAPI](https://img.shields.io/badge/OpenAPI-3.1-blue)](/api/openapi.json)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-api-libs/toggl)](https://goreportcard.com/report/github.com/go-api-libs/toggl)
![Code Coverage](https://img.shields.io/badge/coverage-29%25-red)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

Public Toggl API.
Note:
We use BasicAuth in a specific way. By the standard you provide `Authentication` header with `base64(user_name:password)` as a `credential`.
In our case it will be `base64(user_name:api_token)`.

## Installation

To install the library, use the following command:

```shell
go get github.com/go-api-libs/toggl/pkg/toggl
```

## Usage

### Example: Me

```go
package main

import (
	"context"

	"github.com/go-api-libs/toggl/pkg/toggl"
)

func main() {
	c, err := toggl.NewClient()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	getMeOkJSONResponse, err := c.GetMe(ctx)
	if err != nil {
		panic(err)
	}

	// Use getMeOkJSONResponse object
}

```

## Additional Information

- [**Go Reference**](https://pkg.go.dev/github.com/go-api-libs/toggl/pkg/toggl): The Go reference documentation for the client package.
- [**OpenAPI Specification**](./api/openapi.json): The OpenAPI 3.1.0 specification.
- [**Go Report Card**](https://goreportcard.com/report/github.com/go-api-libs/toggl): Check the code quality report.

## Contributing

If you have any contributions to make, please submit a pull request or open an issue on the [GitHub repository](https://github.com/go-api-libs/toggl).

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
