# Anthropic SDK for Go

> [!NOTE]\
> This is not an official SDK, and I am not affiliated with [Anthropic].\
> Official SDKs from Anthropic are available
> for [Python](https://github.com/anthropics/anthropic-sdk-python),
> and [TypeScript/JavaScript](https://github.com/anthropics/anthropic-sdk-typescript).

Client library for interacting with the [Anthropic] safety-first language model
REST APIs.

## Getting started

### Requirements

- [Go] 1.22+

### Installation and usage

```go
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/unfunco/anthropic-sdk-go"
)

func main() {
    ctx := context.Background()

    transport := &anthropic.Transport{APIKey: os.Getenv("ANTHROPIC_API_KEY")}
    claude := anthropic.NewClient(transport.Client())

    if resp, httpResp, err := claude.Messages.Create(ctx, &anthropic.CreateMessageInput{
        MaxTokens: 1024,
        Messages: []anthropic.Message{
            {Content: "Hello, Claude!", Role: "user"},
        },
        Model: anthropic.Claude3Opus20240229,
    }); err == nil {
        fmt.Println(httpResp.StatusCode)
        fmt.Println(resp)
    } else {
        _, _ = fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
```

### Development and testing

Clone the repository and change into the `anthropic-sdk-go` directory.

```bash
git clone git@github.com:unfunco/anthropic-sdk-go.git
cd anthropic-sdk-go
```

## References

- [Anthropic API](https://www.anthropic.com/api)

## License

© 2024 [Daniel Morris]\
Made available under the terms of the [MIT License].

[anthropic]: https://www.anthropic.com
[daniel morris]: https://unfun.co
[go]: https://go.dev
[mit license]: LICENSE.md
