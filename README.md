# Wire usage example
A simple example of https://github.com/google/wire usage.
More about the tool:
https://go.dev/blog/wire

# Generate
Wire input and generated file is in `internal/wire-gen`.

To generate:
```sh
cd internal/wire-gen
wire
# to show dependencies
wire show
```

# Run

```sh
go run main.go
```
