# Fred

`fred` is a simple wrapper around the
[Fred St. Louis Fed](https://fred.stlouisfed.org) api. 

## Installation

```bash
go get github.com/chrisswanson/fred
```
## Quick Example

This wrapper is designed to be simple to use. Here is a quick example of how to
retrieve a list of sources from the Fred API.

```go
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chrisswanson/fred/fred"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env file")
	}

	fredApiKey := os.Getenv("FRED_API_KEY")

	client, err := fred.NewFredClient(
		fredApiKey,
		1*time.Minute,
		120,
		&zerolog.Logger{},
	)
	if err != nil {
		panic(err)
	}

	params := make(map[string]interface{})

	results, err := client.GetSources(params)
	if err != nil {
		panic(err)
	}

	fmt.Println(len(results.Sources))

	for _, source := range results.Sources {
		fmt.Println(source.ID, source.Name, source.Link)
	}

}```