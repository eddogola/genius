# kanye

A simple wrapper for the API [Kanye Rest](https://kanye.rest).

## Installation

```bash
go get -u github.com/eddogola/kanye
```

## Usage

The API isn't authenticated so you just make requests.

```go
package main

import (
    "fmt"
    "github.com/eddogola/kanye"
)

func main() {
    c := kanye.NewClient() // Initialize a default client
    q, err := c.GetQuote()
    // handle error

    fmt.Println(q)
}

```
