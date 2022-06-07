# Go: Code Generation VS Generics

Simple example of stack implementation with code generation and generics.

## Usage

### Code Generation

Example of generation of stack for type `int`:

```shell
go run stack-generator.go -o stack-int.go -p main -n Int -t int
```

See `go run stack-generator.go --help` for usage details.

### Generic Stack

Example of generic stack usage:

```go
s := Stack[int]{}
```
