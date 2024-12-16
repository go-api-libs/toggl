# Toggl API
[![Go Reference](https://pkg.go.dev/badge/github.com/go-api-libs/toggl.svg)](https://pkg.go.dev/github.com/go-api-libs/toggl/pkg/toggl)
[![OpenAPI](https://img.shields.io/badge/OpenAPI-3.1-blue)](/api/openapi.json)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-api-libs/toggl)](https://goreportcard.com/report/github.com/go-api-libs/toggl)
![Code Coverage](https://img.shields.io/badge/coverage-57%25-yellow)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

The Toggl API allows developers to interact with Toggl's time tracking service programmatically. It provides endpoints for managing time entries, projects, clients, tags, and user information. With this API, you can automate time tracking, generate detailed reports, and integrate Toggl with other tools and services.

## Installation

To install the library, use the following command:

```shell
go get github.com/go-api-libs/toggl/pkg/toggl
```

## Usage

### Example 1: Return Current User

```go
package main

import (
	"context"
	"os"

	"github.com/go-api-libs/toggl/pkg/toggl"
)

func main() {
	c, err := toggl.NewClient("myUsername", os.Getenv("TOGGL_PASSWORD"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	userWithRelated, err := c.GetMe(ctx, &toggl.GetMeParams{WithRelatedData: true})
	if err != nil {
		panic(err)
	}

	// Use userWithRelated object
}

```

### Example 2: Create a time entry

```go
package main

import (
	"context"
	"os"
	"time"

	"github.com/go-api-libs/toggl/pkg/toggl"
)

func main() {
	c, err := toggl.NewClient("myUsername", os.Getenv("TOGGL_PASSWORD"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	timeEntry, err := c.CreateTimeEntry(ctx, 2230580, toggl.NewTimeEntry{
		CreatedWith: "github.com/go-api-libs/toggl",
		Start:       time.Date(2024, time.December, 15, 21, 17, 59, 593648000, time.Local),
		WorkspaceID: 2230580,
	})
	if err != nil {
		panic(err)
	}

	// Use timeEntry object
}

```

### Example 3: Get current time entry

```go
package main

import (
	"context"
	"os"

	"github.com/go-api-libs/toggl/pkg/toggl"
)

func main() {
	c, err := toggl.NewClient("myUsername", os.Getenv("TOGGL_PASSWORD"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	timeEntry, err := c.GetCurrentTimeEntry(ctx)
	if err != nil {
		panic(err)
	}

	// Use timeEntry object
}

```

### Example 4: Stop an existing time entry

```go
package main

import (
	"context"
	"os"

	"github.com/go-api-libs/toggl/pkg/toggl"
)

func main() {
	c, err := toggl.NewClient("myUsername", os.Getenv("TOGGL_PASSWORD"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	timeEntry, err := c.StopTimeEntry(ctx, 2230580, 3730303299)
	if err != nil {
		panic(err)
	}

	// Use timeEntry object
}

```

### Example 5: List time entries

```go
package main

import (
	"context"
	"os"

	"github.com/go-api-libs/toggl/pkg/toggl"
)

func main() {
	c, err := toggl.NewClient("myUsername", os.Getenv("TOGGL_PASSWORD"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	err := c.GetTimeEntries(ctx, &toggl.GetTimeEntriesParams{
		EndDate:   "1984-03-12",
		StartDate: "1984-03-10",
	})
	if err != nil {
		panic(err)
	}
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
