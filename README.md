# dotn

Reflection package to set struct, map, array or slice values via a dot notation syntax string.  
**Warning:** This package only works on **string** paths and values.

## Example

```go
package main

import (
    "github.com/SherClockHolmes/dotn"
)

type Nested struct {
    Slice []string `json:"slice"`
}

type Root struct {
    Nested map[string]*Nested `json:"nested"`
}

func main() {
    root := &Root{
        Nested: map[string]*Nested{
            "key1": &Nested{
                Slice: []string{"one", "two"},
            },
        },
    }

    // Set value of nested slice
    if err := dotn.Set("nested.key1.slice.0", "three", root); err != nil {
        // TODO: Handle error
    }
}
```

## Development

1. Install [Go 1.12+](https://golang.org/)
3. `go test`
