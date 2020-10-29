# vibezy
![license](https://img.shields.io/github/license/dwalker-va/vibezy)

![status](https://github.com/dwalker-va/vibezy/workflows/Go/badge.svg)

A simple Golang SDK for the [OfficeVibe API](https://api.officevibe.com/docs). 

This is not maintained by OfficeVibe. 

## Installing
Use `go get` to install the vibezy Golang module:

`go get github.com/dwalker-va/vibezy`

Include vibezy in your application:

`import "github.com/dwalker-va/vibezy"`

## Usage
Vibezy offers an RPC Client with Request and Response structs to give you full control over how you interact with OfficeVibe's API.

### Instantiate your client
```go
client := vibezy.NewClient("your_officevibe_api_key")
```

### Make requests
```go
resp, err := client.Ping(ctx)
if err != nil { ... }
```

### Make requests with bodies
```go
resp, err := client.CreateGroup(ctx, vibezy.CreateGroupRequest{Name: "new_group"})
if err != nil { ... }
```

## Design
All API calls accept a `context.Context` and support context control, such as cancellation and timeouts.

## Versioning
This project follows [Semantic Versioning 2.0.0](http://semver.org/) guidelines.

## Contributing
Check out the [roadmap](https://github.com/dwalker-va/vibezy/issues/3)

Want to contribute? [see CONTRIBUTING.md](CONTRIBUTING.md)
