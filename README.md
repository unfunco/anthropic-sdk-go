# Anthropic SDK for Go

[![CI](https://github.com/unfunco/anthropic-sdk-go/actions/workflows/ci.yaml/badge.svg)](https://github.com/unfunco/anthropic-sdk-go/actions/workflows/ci.yaml)
[![License: MIT](https://img.shields.io/badge/License-MIT-purple.svg)](https://opensource.org/licenses/MIT)

> [!NOTE]\
> This is not an official SDK, and I am not affiliated with [Anthropic].\
> Official SDKs from Anthropic are available
> for [Python](https://github.com/anthropics/anthropic-sdk-python),
> and [TypeScript/JavaScript](https://github.com/anthropics/anthropic-sdk-typescript).

Client library for interacting with the [Anthropic] safety-first language model
REST APIs.\
Documentation for the REST API can be found at [docs.anthropic.com].

## Getting started

### Requirements

- [Claude API key]
- [Go] 1.22+

### Installation and usage

Import the module and run the `go get` command without any arguments to resolve
and add the SDK to your dependencies.

```go
import "github.com/unfunco/anthropic-sdk-go"
```

Construct a new REST API client with a `http.Client` derived from a Transport
containing your [Claude API key]. The derived HTTP client will automatically add
the API key as a header to each request sent to the API.

```go
transport := anthropic.NewTransport(os.Getenv("ANTHROPIC_API_KEY"))
claude := anthropic.NewClient(transport.Client())
```

Once constructed, you can use the client to interact with the REST API.

```go
data, _, err := claude.Messages.Create(
    context.Background(),
    &anthropic.CreateMessageInput{
        MaxTokens: 1024,
        Messages: []anthropic.Message{
            {
                Content: "Hello, Claude!",
                Role:    "user",
            },
        },
        Model: anthropic.Claude3Opus20240229,
    },
)
```

### Development and testing

Clone the repository and change into the `anthropic-sdk-go` directory.

```bash
git clone git@github.com:unfunco/anthropic-sdk-go.git
cd anthropic-sdk-go
```

Run the unit tests with coverage to ensure everything is working as expected.

```bash
go test -cover -v ./...
```

## References

- [Anthropic API](https://www.anthropic.com/api)

## License

© 2024 [Daniel Morris]\
Made available under the terms of the [MIT License].

[anthropic]: https://www.anthropic.com
[claude api key]: https://www.anthropic.com/api
[daniel morris]: https://unfun.co
[docs.anthropic.com]: https://docs.anthropic.com
[go]: https://go.dev
[mit license]: LICENSE.md
