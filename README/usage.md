To install the library, use the following command:

```shell
go get github.com/go-api-libs/toggl/pkg/toggl
```


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
		Start:       mustParseTime("2024-12-15T21:17:59.593648+01:00"),
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
	timeEntries, err := c.ListTimeEntries(ctx, &toggl.ListTimeEntriesParams{
		Before:         mustParseTime("2024-12-16T03:25:20+01:00"),
		EndDate:        mustParseTime("2024-12-16T03:25:20+01:00"),
		IncludeSharing: true,
		Meta:           true,
		Since:          1734304527,
		StartDate:      mustParseTime("2024-12-16T03:25:20+01:00"),
	})
	if err != nil {
		panic(err)
	}

	// Use timeEntries slice
}

```

### Example 6: Create a new organization

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
	simpleOrganization, err := c.CreateOrganization(ctx, toggl.NewOrganization{
		Name:          "Your Organization",
		WorkspaceName: "Your Workspace",
	})
	if err != nil {
		panic(err)
	}

	// Use simpleOrganization object
}

```

### Example 7: List my organizations

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
	organizations, err := c.ListOrganizations(ctx)
	if err != nil {
		panic(err)
	}

	// Use organizations slice
}

```

### Example 8: Get organization data

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
	organization, err := c.GetOrganization(ctx, 9011051)
	if err != nil {
		panic(err)
	}

	// Use organization object
}

```
